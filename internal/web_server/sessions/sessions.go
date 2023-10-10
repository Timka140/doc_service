package sessions

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var Ses ISessions

type ISessions interface {
	GetSes(c *gin.Context) ISession
	Add(key string, ses ISession) error
	Get(key string) (ISession, error)
	RangeSes(fn func(ses ISession))
}
type TSessions struct {
	sessions sync.Map
}

// func NewSessions() (ISessions, error) {
// 	t :=
// 	return t, nil
// }

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
		return nil, fmt.Errorf("TSessions.Get(): сессия не найдена") //nil
	}

	ses, ok := v.(ISession)
	if !ok {
		return nil, fmt.Errorf("TSessions.Get(): неизвестный тип данных")
	}

	return ses, nil
}

func (t *TSessions) GetSes(c *gin.Context) ISession {
	val, err := c.Cookie("AccessToken")
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/gui/login")
		// redirect.Redirect(c, "/gui/login")
		return nil
		// return nil, fmt.Errorf(" GetSes(): чтение cookie, err=%w", err)
	}

	ses, err := t.Get(val)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/gui/login")
		// redirect.Redirect(c, "/gui/login")
		return nil
		// return nil, fmt.Errorf(" GetSes(): чтение сессии, err=%w", err)
	}

	if ses == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/gui/login")
		// redirect.Redirect(c, "/gui/login")
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
