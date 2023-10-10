package socket

import (
	"log"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/google/uuid"
)

type TCheckSession struct {
	data map[string]interface{}
	ses  sessions.ISession
	pid  string
}

func newCheckSessionSocket(in *TSocketValue) (ISocket, error) {
	t := &TCheckSession{
		data: in.Data,
		ses:  in.Ses,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TCheckSession) Start() error {
	path, ok := t.data["path"].(string)
	if ok {
		t.ses.SetCurrentPage(path)
	}
	return nil
}

func (t *TCheckSession) GetPid() string {
	return t.pid
}

func (t *TCheckSession) Response() (map[string]interface{}, error) {
	t.data["login"] = t.ses.Authorization()
	return t.data, nil
}

func (t *TCheckSession) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("CheckSession", newCheckSessionSocket)
	if err != nil {
		log.Printf("CheckSession(): не удалось добавить в конструктор")
	}
}
