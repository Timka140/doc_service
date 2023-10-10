package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/google/uuid"
)

type TTaskList struct {
	data map[string]interface{}
	ses  sessions.ISession
	pid  string

	path string
}

func newTaskListSocket(in *TSocketValue) (ISocket, error) {
	t := &TTaskList{
		data: in.Data,
		ses:  in.Ses,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TTaskList) Start() error {
	var err error
	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TTaskList.Start(): не прочитано исполнение")
	}

	t.path, _ = t.data["path"].(string)
	if t.path == "" {
		t.path = "/"
	}

	switch execution {
	case "list":
		err = t.list()
		if err != nil {
			return fmt.Errorf("TaskList.Start(): получение списка, err=%w", err)
		}
	case "open":
		err = t.open()
		if err != nil {
			return fmt.Errorf("TaskList.Start(): открытие, err=%w", err)
		}

	}
	return nil
}

func (t *TTaskList) GetPid() string {
	return t.pid
}

func (t *TTaskList) list() error {
	return nil
}

func (t *TTaskList) open() error {
	tp_task, ok := t.data["tp_task"].(float64)
	if !ok {
		return fmt.Errorf("TTaskList.Start(): не прочитано исполнение")
	}

	name, ok := t.data["name"].(string)
	if !ok {
		return fmt.Errorf("TTaskList.Start(): не прочитано название")
	}

	if t.path == "/" {
		t.path = ""
	}
	switch tp_task {
	case 1: //База данных
		t.path = fmt.Sprintf("%v/%v", t.path, name)
	case 2: //Каталог
		t.path = fmt.Sprintf("%v/%v", t.path, name)
	case 3: // Задача
	}
	return nil
}
func (t *TTaskList) Response() (map[string]interface{}, error) {
	var err error
	var rows []db.Tasks
	err = db.DB.Table("tasks").Select("id, path, tp, name, comment").Where("path LIKE ?", t.path).Scan(&rows).Error
	if err != nil {
		return nil, fmt.Errorf("TaskList.Start(): чтение таблицы, err=%w", err)
	}

	t.data["tasks"] = rows
	t.data["path"] = t.path
	return t.data, nil
}

func (t *TTaskList) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("TaskList", newTaskListSocket)
	if err != nil {
		log.Printf("TaskList(): не удалось добавить в конструктор")
	}
}
