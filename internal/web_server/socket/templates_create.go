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
			return fmt.Errorf("TaskCreate.Start(): создание шаблона, err=%w", err)
		}
	case "remove":
		err = t.remove()
		if err != nil {
			return fmt.Errorf("TaskCreate.Start(): удаление, err=%w", err)
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
		Name:   name,
		Tp:     tp_temp,
		Path:   ph,
		UserID: t.ses.ID(),
	}

	err := db.DB.Table("templates").Create(&rows).Error
	if err != nil {
		return fmt.Errorf("TaskCreate.Start(): создание записи, err=%w", err)
	}

	return nil
}

func (t *TTemplatesCreate) remove() error {
	temps, ok := t.data["temps"].([]interface{})
	if !ok {
		return fmt.Errorf("TTemplatesCreate.remove(): отсутствуют данные")
	}

	for _, item := range temps {
		v, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		sl, ok := v["select"].(bool)
		if !ok {
			continue
		}

		if !sl {
			continue
		}

		tp, ok := v["Tp"].(float64)
		if !ok {
			continue
		}

		id, ok := v["Id"].(float64)
		if !ok {
			continue
		}

		if tp == 3 {
			err := db.DB.Exec("DELETE FROM templates WHERE id = ?", id).Error
			if err != nil {
				return fmt.Errorf("TaskCreate.Start(): удаление, err=%w", err)
			}
			continue
		}

		path, ok := v["Path"].(string)
		if !ok {
			continue
		}
		name, ok := v["Name"].(string)
		if !ok {
			continue
		}

		if path == "/" {
			path = ""
		}

		path = fmt.Sprintf("%v/%v", path, name) + "%"
		err := db.DB.Exec("DELETE FROM templates WHERE path LIKE ? OR id = ?", path, id).Error
		if err != nil {
			return fmt.Errorf("TaskCreate.Start(): удаление, err=%w", err)
		}
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
		Name:   name,
		Tp:     tp_temp,
		Path:   ph,
		UserID: t.ses.ID(),
	}
	err := db.DB.Table("templates").Create(&row).Error
	if err != nil {
		return fmt.Errorf("TaskCreate.Start(): создание записи, err=%w", err)
	}

	// taskName := fmt.Sprintf("template_%v", row.Id)
	// task := db.Template{
	// 	Name:     taskName,
	// 	TaskID:   row.Id,
	// 	PathBase: fmt.Sprintf("templates/%v", taskName),
	// }
	// err = db.DB.Table("template").Create(&task).Error
	// if err != nil {
	// 	return fmt.Errorf("TaskCreate.Start(): создание записи task, err=%w", err)
	// }

	return nil
}

func init() {
	err := constructors.Add("TemplatesCreate", newTemplatesCreateSocket)
	if err != nil {
		log.Printf("TaskCreate(): не удалось добавить в конструктор")
	}
}
