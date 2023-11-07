package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/template"
	"projects/doc/doc_service/internal/web_server/sessions"
	"projects/doc/doc_service/pkg/types"

	"github.com/google/uuid"
)

type TTemplate struct {
	data        map[string]interface{}
	ses         sessions.ISession
	pid         string
	template_id string

	tmp types.ITemplate

	path string
}

func newTemplateSocket(in *TSocketValue) (ISocket, error) {
	t := &TTemplate{
		data: in.Data,
		ses:  in.Ses,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TTemplate) Start() error {
	var err error
	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TTemplate.Start(): не прочитано исполнение")
	}

	tId, ok := t.data["template_id"]
	if !ok {
		return fmt.Errorf("TTemplate.Start(): template_id не задан")
	}
	t.template_id = fmt.Sprintf("%v", tId)

	switch execution {
	case "init":
		err = t.init()
		if err != nil {
			return fmt.Errorf("TTemplate.Start(): инициализация шаблона, err=%w", err)
		}
	case "open":
	}
	return nil
}

func (t *TTemplate) GetPid() string {
	return t.pid
}

func (t *TTemplate) Response() (map[string]interface{}, error) {
	t.data["path"] = t.path
	return t.data, nil
}

func (t *TTemplate) Stop() error {

	return nil
}

func (t *TTemplate) init() error {
	var err error
	t.tmp, err = template.New(t.template_id)
	if err != nil {
		return fmt.Errorf("TTemplate.Start(): инициализация шаблона, err=%w", err)
	}

	t.data["load_file"] = t.tmp.IsFile()
	t.data["file_name"] = t.tmp.Name()
	t.data["tm_update"] = t.tmp.UpdateTime()

	return nil
}

func init() {
	err := constructors.Add("Template", newTemplateSocket)
	if err != nil {
		log.Printf("Template(): не удалось добавить в конструктор")
	}
}
