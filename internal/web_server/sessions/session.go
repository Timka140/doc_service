package sessions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	"projects/doc/doc_service/internal/db"

	"github.com/gorilla/websocket"
)

type ISession interface {
	Authorization() bool
	GetLogin() string
	CherRights(in int) bool
	GetConn() *TConn
	SendMessage(params map[string]interface{}) (err error)

	GetCurrentPage() string
	SetCurrentPage(string)
}

type TConn struct {
	sync.Mutex
	*websocket.Conn //Веб сокет для ответов из потока
}

type TSession struct {
	authorization bool
	login         string
	rights        TRights
	conn          TConn

	currentPage string
}

func NewSession(hash string) (ISession, error) {
	t := &TSession{}

	var buf []byte
	var user db.Users
	err := db.DB.Table("users").Select("id, login, rights").Where("password = ? AND state = 1", hash).First(&user).Error
	// err := db.DB.QueryRow("SELECT id, login, rights FROM users WHERE password = ? AND state = 1", hash).Scan(&t.id, &t.login, &buf)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
	default:
		return nil, fmt.Errorf("NewSession(): чтение сессии, err=%w", err)
	}

	if len(buf) > 0 {
		err = json.Unmarshal(buf, &t.rights)
		if err != nil {
			return nil, fmt.Errorf("NewSession(): распаковка прав, err=%w", err)
		}
	}

	t.authorization = true

	return t, nil
}

func (t *TSession) GetLogin() string {
	return t.login
}

func (t *TSession) CherRights(in int) bool {
	for _, v := range t.rights.Rights {
		if v == in {
			return true
		}
	}
	return false
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
		return
	}

	t.conn.Lock()
	err = t.conn.WriteMessage(websocket.TextMessage, send)
	if err != nil {
		return
	}
	t.conn.Unlock()

	return err
}

func (t *TSession) GetCurrentPage() string {
	return t.currentPage
}

func (t *TSession) SetCurrentPage(in string) {
	t.currentPage = in
}
