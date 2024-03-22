package grpc_sessions

import (
	"fmt"
	"sync"
	"time"
)

var Ses ISessions

type ISessions interface {
	GetSes(token string) ISession
	Add(key string, ses ISession) error
	Get(key string) (ISession, error)
	RangeSes(fn func(ses ISession))
	Delete(key string)
}
type TSessions struct {
	sessions sync.Map
}

func (t *TSessions) Add(key string, ses ISession) error {
	if key == "" {
		return fmt.Errorf("TSessions.Add(): ключ не задан")
	}

	if ses == nil {
		return fmt.Errorf("TSessions.Add(): сессия не задана")
	}

	t.sessions.Store(key, ses)
	return nil
}

func (t *TSessions) Get(key string) (ISession, error) {
	if key == "" {
		return nil, fmt.Errorf("TSessions.Get(): ключ не задан")
	}

	v, ok := t.sessions.Load(key)
	if !ok {
		time.Sleep(2 * time.Second)
		return nil, fmt.Errorf("TSessions.Get(): сессия не найдена") //nil
	}

	ses, ok := v.(ISession)
	if !ok {
		return nil, fmt.Errorf("TSessions.Get(): неизвестный тип данных")
	}

	return ses, nil
}

func (t *TSessions) Delete(key string) {
	t.sessions.Delete(key)
}

func (t *TSessions) GetSes(token string) ISession {
	var err error
	ses, err := t.Get(token)
	if err != nil {
		return nil
	}

	if ses == nil {
		return nil
	}

	return ses
}

func (t *TSessions) RangeSes(fn func(ses ISession)) {
	t.sessions.Range(func(key, value any) bool {
		ses, ok := value.(ISession)
		if !ok {
			return true
		}
		fn(ses)
		return true
	})
}

func init() {
	t := &TSessions{}
	Ses = t
}
