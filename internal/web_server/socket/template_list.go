package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/google/uuid"
)

type TTemplateList struct {
	data map[string]interface{}
	ses  sessions.ISession
	pid  string

	path string
}

func newTemplateListSocket(in *TSocketValue) (ISocket, error) {
	t := &TTemplateList{
		data: in.Data,
		ses:  in.Ses,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TTemplateList) Start() error {
	var err error
	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TTemplateList.Start(): не прочитано исполнение")
	}

	t.path, _ = t.data["path"].(string)
	if t.path == "" {
		t.path = "/"
	}

	switch execution {
	case "list":
		err = t.list()
		if err != nil {
			return fmt.Errorf("TemplateList.Start(): получение списка, err=%w", err)
		}
	case "open":
		err = t.open()
		if err != nil {
			return fmt.Errorf("TemplateList.Start(): открытие, err=%w", err)
		}
	}
	return nil
}

func (t *TTemplateList) GetPid() string {
	return t.pid
}

func (t *TTemplateList) list() error {
	return nil
}

func (t *TTemplateList) open() error {
	tp_temp, ok := t.data["tp_temp"].(float64)
	if !ok {
		return fmt.Errorf("TTemplateList.Start(): не прочитано исполнение")
	}

	name, ok := t.data["name"].(string)
	if !ok {
		return fmt.Errorf("TTemplateList.Start(): не прочитано название")
	}

	if t.path == "/" {
		t.path = ""
	}
	switch tp_temp {
	case 1: //База данных
		t.path = fmt.Sprintf("%v/%v", t.path, name)
	case 2: //Каталог
		t.path = fmt.Sprintf("%v/%v", t.path, name)
	case 3: // Задача
	}
	return nil
}
func (t *TTemplateList) Response() (map[string]interface{}, error) {
	var err error
	var rows []db.Tasks

	tx := db.DB.Table("templates").Select("id, path, tp, name, comment")
	if t.ses.Rights([]int{sessions.CAdministrator}) {
		tx = tx.Where("path LIKE ?", t.path)
	} else {
		tx = tx.Where("path LIKE ? AND user_id = ?", t.path, t.ses.ID())
	}

	err = tx.Scan(&rows).Error
	if err != nil {
		return nil, fmt.Errorf("TemplateList.Start(): чтение таблицы, err=%w", err)
	}

	// err = db.DB.Table("templates").Select("id, path, tp, name, comment").Where("path LIKE ?", t.path).Scan(&rows).Error
	// if err != nil {
	// 	return nil, fmt.Errorf("TemplateList.Start(): чтение таблицы, err=%w", err)
	// }

	t.data["temps"] = rows
	t.data["path"] = t.path
	return t.data, nil
}

func (t *TTemplateList) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("TemplateList", newTemplateListSocket)
	if err != nil {
		log.Printf("TemplateList(): не удалось добавить в конструктор")
	}
}
