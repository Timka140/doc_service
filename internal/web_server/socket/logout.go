package socket

import (
	"log"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/google/uuid"
)

type TLogout struct {
	data map[string]interface{}
	pid  string
	ses  sessions.ISession
}

func newLogout(in *TSocketValue) (ISocket, error) {
	t := &TLogout{
		data: in.Data,
		pid:  uuid.NewString(),
		ses:  in.Ses,
	}

	return t, nil
}

func (t *TLogout) Start() error {
	if !t.ses.Authorization() {
		return nil
	}
	token := t.ses.Token()
	sessions.Ses.Delete(token)
	return nil
}

func (t *TLogout) GetPid() string {
	return t.pid
}

func (t *TLogout) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TLogout) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("Logout", newLogout)
	if err != nil {
		log.Printf("Logout(): не удалось добавить в конструктор")
	}
}
