package sessions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"projects/doc/doc_service/internal/db"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ISession interface {
	ID() int64
	Authorization() bool
	GetLogin() string
	Rights(in []int) bool
	RightsVue() map[string]interface{}
	GetConn() *TConn
	SendMessage(params map[string]interface{}) (err error)

	GetCurrentPage() string
	SetCurrentPage(string)

	Token() string
	CloseSocket()
	OpenSocket()

	Online()
}

type TConn struct {
	sync.Mutex
	*websocket.Conn //Веб сокет для ответов из потока
}

type tOnline struct {
	sync.Mutex
	online     bool
	onlineTime *time.Time
	close      time.Time
}

type TSession struct {
	id    int64
	token string
	// Время закрытия сессии
	update time.Time
	online tOnline

	authorization bool
	login         string
	rights        TRights
	conn          TConn

	currentPage string
}

func NewSession(hash string) (ISession, error) {
	t := &TSession{
		token:  uuid.NewString(),
		update: time.Now(),
	}

	var user db.Users
	err := db.DB.Table("users").Select("id, login, rights").Where("password = ? AND state = 1", hash).First(&user).Error
	// err := db.DB.QueryRow("SELECT id, login, rights FROM users WHERE password = ? AND state = 1", hash).Scan(&t.id, &t.login, &buf)
	switch err {
	case sql.ErrNoRows:
		time.Sleep(2 * time.Second)
		return nil, nil
	case nil:
	default:
		return nil, fmt.Errorf("NewSession(): чтение сессии, err=%w", err)
	}

	t.id = user.Id
	t.login = user.Login
	t.rights.SetDB(user.Rights)
	t.authorization = true

	go t.monitor()

	return t, nil
}

// monitor -- проверка онлайна пользователя
func (t *TSession) monitor() {
	check := func() bool {
		t.online.Lock()
		defer t.online.Unlock()

		if !t.online.online {
			duration := t.update.Sub(t.online.close)
			if duration.Minutes() > 5 {
				//удаляем сессию
				Ses.Delete(t.token)
				return true
			}
		}

		// Если авторизация через api смотрю последний online
		if t.online.onlineTime != nil {
			duration := t.update.Sub(*t.online.onlineTime)
			if duration.Minutes() > 5 {
				t.online.online = false
				t.online.close = *t.online.onlineTime
				Ses.Delete(t.token)
				return true
			}
		}
		return false
	}
	for {
		time.Sleep(time.Second * 30)
		if check() {
			return
		}
		t.update = time.Now()

	}
}

func (t *TSession) ID() int64 {
	return t.id
}

func (t *TSession) OpenSocket() {
	t.online.online = true
}

func (t *TSession) Token() string {
	return t.token
}

func (t *TSession) GetLogin() string {
	return t.login
}

func (t *TSession) Rights(in []int) bool {
	return t.rights.Check(in)
}
func (t *TSession) RightsVue() map[string]interface{} {
	return t.rights.Vue()
}

func (t *TSession) Online() {
	t.online.Lock()
	tm := time.Now()
	t.online.online = true
	t.online.onlineTime = &tm
	t.online.Unlock()
}

func (t *TSession) CloseSocket() {
	t.online.Lock()
	t.online.online = false
	t.online.close = time.Now()
	t.online.Unlock()
}

func (t *TSession) Authorization() bool {
	return t.authorization
}

func (t *TSession) GetConn() *TConn {
	return &t.conn
}

func (t *TSession) SendMessage(params map[string]interface{}) (err error) {
	send, err := json.Marshal(params)
	if err != nil {
		return err
	}

	if t.conn.Conn == nil {
		return nil
	}

	t.conn.Lock()
	defer t.conn.Unlock()
	err = t.conn.WriteMessage(websocket.TextMessage, send)
	if err != nil {
		return err
	}

	return nil
}

func (t *TSession) GetCurrentPage() string {
	return t.currentPage
}

func (t *TSession) SetCurrentPage(in string) {
	t.currentPage = in
}
