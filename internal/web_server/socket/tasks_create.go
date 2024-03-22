package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/google/uuid"
)

type TTaskCreate struct {
	data map[string]interface{}
	ses  sessions.ISession
	pid  string
}

func newTaskCreateSocket(in *TSocketValue) (ISocket, error) {
	t := &TTaskCreate{
		data: in.Data,
		ses:  in.Ses,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TTaskCreate) Start() error {
	if !t.ses.Rights([]int{sessions.CGuest}) {
		return nil
	}
	var err error
	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TListDocxServices.Start(): не прочитано исполнение")
	}

	switch execution {

	case "store":
		t.data["tp_task"] = 1
	case "catalog":
		t.data["tp_task"] = 2
		err = t.createCatalog()
		if err != nil {
			return fmt.Errorf("TaskCreate.Start(): создание каталога, err=%w", err)
		}
	case "tasks":
		t.data["tp_task"] = 3
		err = t.createTask()
		if err != nil {
			return fmt.Errorf("TaskCreate.Start(): создание задачи, err=%w", err)
		}
	}

	return nil
}

func (t *TTaskCreate) GetPid() string {
	return t.pid
}

func (t *TTaskCreate) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TTaskCreate) Stop() error {
	return nil
}

func (t *TTaskCreate) createCatalog() error {
	ph, ok := t.data["path"].(string)
	if !ok {
		return fmt.Errorf("TTaskCreate.createCatalog(): путь не указан")
	}

	data, ok := t.data["data"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TTaskCreate.createCatalog(): отсутствуют данные")
	}
	tp_task, ok := t.data["tp_task"].(int)
	if !ok {
		return fmt.Errorf("TTaskCreate.createCatalog(): тип не задан")
	}
	name, ok := data["name"].(string)
	if !ok {
		return fmt.Errorf("TTaskCreate.createCatalog(): имя не задано")
	}
	rows := db.Tasks{
		Name: name,
		Tp:   tp_task,
		Path: ph,
	}
	err := db.DB.Table("tasks").Create(&rows).Error
	if err != nil {
		return fmt.Errorf("TaskCreate.Start(): создание записи, err=%w", err)
	}

	return nil
}

func (t *TTaskCreate) createTask() error {
	ph, ok := t.data["path"].(string)
	if !ok {
		return fmt.Errorf("TTaskCreate.createTask(): путь не указан")
	}

	data, ok := t.data["data"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TTaskCreate.createTask(): отсутствуют данные")
	}
	tp_task, ok := t.data["tp_task"].(int)
	if !ok {
		return fmt.Errorf("TTaskCreate.createTask(): тип не задан")
	}
	name, ok := data["name"].(string)
	if !ok {
		return fmt.Errorf("TTaskCreate.createTask(): имя не задано")
	}
	row := db.Tasks{
		Name: name,
		Tp:   tp_task,
		Path: ph,
	}
	err := db.DB.Table("tasks").Create(&row).Error
	if err != nil {
		return fmt.Errorf("TaskCreate.Start(): создание записи, err=%w", err)
	}

	taskName := fmt.Sprintf("task_%v", row.Id)
	task := db.Task{
		Name:     taskName,
		TaskID:   row.Id,
		PathBase: fmt.Sprintf("tasks/%v", taskName),
	}
	err = db.DB.Table("task").Create(&task).Error
	if err != nil {
		return fmt.Errorf("TaskCreate.Start(): создание записи task, err=%w", err)
	}

	return nil
}

func init() {
	err := constructors.Add("TaskCreate", newTaskCreateSocket)
	if err != nil {
		log.Printf("TaskCreate(): не удалось добавить в конструктор")
	}
}
