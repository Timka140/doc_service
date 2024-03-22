package grpc_sessions

import (
	"database/sql"
	"fmt"
	"projects/doc/doc_service/internal/db"
	"sync"
	"time"

	"github.com/google/uuid"
)

type ISession interface {
	ID() int64
	Authorization() bool
	Rights(in []int) bool
	Online() bool
}

type tOnline struct {
	sync.Mutex
	online     bool
	onlineTime *time.Time
	close      time.Time
}

type TSession struct {
	id            int64
	token         string
	name          string
	update        time.Time
	status        tOnline
	authorization bool

	rights TRights
}

func NewSession(key string) (ISession, error) {
	t := &TSession{
		token:  uuid.NewString(),
		update: time.Now(),
	}

	var service db.Services
	err := db.DB.Table("services").Select("id, name, rights").Where("key = ? AND state = 1", key).First(&service).Error
	// err := db.DB.QueryRow("SELECT id, login, rights FROM users WHERE password = ? AND state = 1", hash).Scan(&t.id, &t.login, &buf)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
	default:
		return nil, fmt.Errorf("NewSession(): чтение сессии, err=%w", err)
	}

	t.id = service.Id
	t.name = service.Name
	t.rights.SetDB(service.Rights)
	t.authorization = true

	return t, nil
}

func (t *TSession) ID() int64 {
	return t.id
}

func (t *TSession) Rights(in []int) bool {
	return t.rights.Check(in)
}

func (t *TSession) Authorization() bool {
	return t.authorization
}

func (t *TSession) Online() bool {
	return t.status.online
}
