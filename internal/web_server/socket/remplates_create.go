package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/google/uuid"
)

type TTemplatesCreate struct {
	data map[string]interface{}
	ses  sessions.ISession
	pid  string
}

func newTemplatesCreateSocket(in *TSocketValue) (ISocket, error) {
	t := &TTemplatesCreate{
		data: in.Data,
		ses:  in.Ses,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TTemplatesCreate) Start() error {
	var err error
	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TListDocxServices.Start(): не прочитано исполнение")
	}

	switch execution {

	case "store":
		t.data["tp_temp"] = 1
	case "catalog":
		t.data["tp_temp"] = 2
		err = t.createCatalog()
		if err != nil {
			return fmt.Errorf("TaskCreate.Start(): создание каталога, err=%w", err)
		}
	case "template":
		t.data["tp_temp"] = 3
		err = t.createTemplate()
		if err != nil {
			return fmt.Errorf("TaskCreate.Start(): создание задачи, err=%w", err)
		}
	}

	return nil
}

func (t *TTemplatesCreate) GetPid() string {
	return t.pid
}

func (t *TTemplatesCreate) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TTemplatesCreate) Stop() error {
	return nil
}

func (t *TTemplatesCreate) createCatalog() error {
	ph, ok := t.data["path"].(string)
	if !ok {
		return fmt.Errorf("TTemplatesCreate.createCatalog(): путь не указан")
	}

	data, ok := t.data["data"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TTemplatesCreate.createCatalog(): отсутствуют данные")
	}
	tp_temp, ok := t.data["tp_temp"].(int)
	if !ok {
		return fmt.Errorf("TTemplatesCreate.createCatalog(): тип не задан")
	}
	name, ok := data["name"].(string)
	if !ok {
		return fmt.Errorf("TTemplatesCreate.createCatalog(): имя не задано")
	}
	rows := db.Templates{
		Name: name,
		Tp:   tp_temp,
		Path: ph,
	}
	err := db.DB.Table("templates").Create(&rows).Error
	if err != nil {
		return fmt.Errorf("TaskCreate.Start(): создание записи, err=%w", err)
	}

	return nil
}

func (t *TTemplatesCreate) createTemplate() error {
	ph, ok := t.data["path"].(string)
	if !ok {
		return fmt.Errorf("TTemplatesCreate.createTemplate(): путь не указан")
	}

	data, ok := t.data["data"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TTemplatesCreate.createTemplate(): отсутствуют данные")
	}
	tp_temp, ok := t.data["tp_temp"].(int)
	if !ok {
		return fmt.Errorf("TTemplatesCreate.createTemplate(): тип не задан")
	}
	name, ok := data["name"].(string)
	if !ok {
		return fmt.Errorf("TTemplatesCreate.createTemplate(): имя не задано")
	}
	row := db.Templates{
		Name: name,
		Tp:   tp_temp,
		Path: ph,
	}
	err := db.DB.Table("templates").Create(&row).Error
	if err != nil {
		return fmt.Errorf("TaskCreate.Start(): создание записи, err=%w", err)
	}

	taskName := fmt.Sprintf("template_%v", row.Id)
	task := db.Template{
		Name:     taskName,
		TaskID:   row.Id,
		PathBase: fmt.Sprintf("templates/%v", taskName),
	}
	err = db.DB.Table("template").Create(&task).Error
	if err != nil {
		return fmt.Errorf("TaskCreate.Start(): создание записи task, err=%w", err)
	}

	return nil
}

func init() {
	err := constructors.Add("TemplatesCreate", newTemplatesCreateSocket)
	if err != nil {
		log.Printf("TaskCreate(): не удалось добавить в конструктор")
	}
}
