package socket

import (
	"database/sql"
	"fmt"
	"log"
	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/internal/grpc_server/grpc_sessions"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/google/uuid"
)

type TSettingsServices struct {
	data map[string]interface{}
	ses  sessions.ISession
	pid  string
}

func newSettingsServicesSocket(in *TSocketValue) (ISocket, error) {
	t := &TSettingsServices{
		data: in.Data,
		ses:  in.Ses,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TSettingsServices) Start() error {
	var err error
	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TSettingsServices.Start(): не прочитано исполнение")
	}

	switch execution {
	case "check_name":
		err = t.check_name()
		if err != nil {
			return fmt.Errorf("TSettingsServices.Start(): проверка имени, err=%w", err)
		}
	case "create_service":
		err = t.create_service()
		if err != nil {
			return fmt.Errorf("TSettingsServices.Start(): создание сервиса, err=%w", err)
		}
	case "update_service":
		err = t.update_service()
		if err != nil {
			return fmt.Errorf("TSettingsServices.Start(): обновление сервиса, err=%w", err)
		}
	case "info_service":
		err = t.info_service()
		if err != nil {
			return fmt.Errorf("TSettingsServices.Start(): получение данных о сервисе, err=%w", err)
		}
	case "remove_service":
		err = t.remove_service()
		if err != nil {
			return fmt.Errorf("TSettingsServices.Start(): удаление сервиса, err=%w", err)
		}
	case "list":
		err = t.list()
		if err != nil {
			return fmt.Errorf("TSettingsServices.Start(): список сервисов, err=%w", err)
		}
	}
	return nil
}

func (t *TSettingsServices) GetPid() string {
	return t.pid
}

func (t *TSettingsServices) check_name() error {
	name, ok := t.data["name"].(string)
	if !ok {
		return fmt.Errorf("TSettingsServices.check_login(): приведение к строке, name")
	}

	var err error
	var row db.Services

	tx := db.DB.Table("services").Select("id")
	if t.ses.Rights([]int{sessions.CAdministrator}) {
		tx.Where("name = ?", name)
	} else {
		tx.Where("name = ? AND user_id", name, t.ses.ID())
	}
	err = tx.Scan(&row).Error
	switch err {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsServices.check_login(): проверка логина, err=%w", err)
	}
	t.data["present_name"] = 2
	if row.Id != 0 {
		t.data["present_name"] = 1
	}

	return nil
}
func (t *TSettingsServices) create_service() error {
	var err error
	var service db.Services

	sv, ok := t.data["service"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsServices.create_service(): не удалось привести к map, err=%w", err)
	}

	service.UserID = t.ses.ID()
	service.Name, ok = sv["name"].(string)
	if !ok || service.Name == "" {
		return fmt.Errorf("TSettingsServices.create_service(): некорректное имя, err=%w", err)
	}

	service.State = 1
	service.Key = uuid.NewString()

	rights, ok := sv["rights"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsServices.create_service(): некорректные права, err=%w", err)
	}

	rg := grpc_sessions.NewRights()

	r, err := rg.Set(rights)
	if err != nil {
		return fmt.Errorf("TSettingsServices.create_service(): установка прав, err=%w", err)
	}
	service.Rights = r.Get()

	service.Comment, ok = sv["comment"].(string)
	if !ok {
		return fmt.Errorf("TSettingsServices.create_service(): некорректный комментарий, err=%w", err)
	}

	err = db.DB.Table("services").Create(&service).Error
	switch err {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsServices.create_service(): добавление сервиса, err=%w", err)
	}
	return nil
}
func (t *TSettingsServices) update_service() error {
	var err error

	sv, ok := t.data["service"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsServices.update_service(): не удалось привести к map, err=%w", err)
	}

	id, ok := sv["id"].(float64)
	if !ok {
		return fmt.Errorf("TSettingsServices.update_service(): id не установлен, err=%w", err)
	}

	name, ok := sv["name"].(string)
	if !ok || name == "" {
		return fmt.Errorf("TSettingsServices.update_service(): некорректное имя, err=%w", err)
	}

	rights, ok := sv["rights"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsServices.update_service(): некорректные права, err=%w", err)
	}

	stateBool, ok := sv["state"].(bool)
	if !ok {
		return fmt.Errorf("TSettingsServices.update_service(): state не установлен, err=%w", err)
	}
	state := 0
	if stateBool {
		state = 1
	}

	rg := grpc_sessions.NewRights()

	r, err := rg.Set(rights)
	if err != nil {
		return fmt.Errorf("TSettingsServices.update_service(): установка прав, err=%w", err)
	}

	comment, ok := sv["comment"].(string)
	if !ok {
		return fmt.Errorf("TSettingsServices.update_service(): некорректный комментарий, err=%w", err)
	}

	service := map[string]interface{}{
		"name":    name,
		"rights":  r.Get(),
		"comment": comment,
		"state":   state,
	}

	tx := db.DB.Table("services").Updates(service)
	if t.ses.Rights([]int{sessions.CAdministrator}) {
		tx.Where("id = ?", id)
	} else {
		tx.Where("id = ? AND user_id", id, t.ses.ID())
	}

	switch tx.Error {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsServices.create_service(): обновления сервиса, err=%w", err)
	}

	return nil
}

func (t *TSettingsServices) info_service() error {
	var err error
	sv, ok := t.data["service"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsServices.update_service(): не удалось привести к map, err=%w", err)
	}

	id, ok := sv["id"].(float64)
	if !ok {
		return fmt.Errorf("TSettingsServices.update_service(): id не установлен, err=%w", err)
	}

	var service db.Services
	tx := db.DB.Table("services").Select("id, name,key,state, rights, comment")
	if t.ses.Rights([]int{sessions.CAdministrator}) {
		tx.Where("id = ?", id)
	} else {
		tx.Where("id = ? AND user_id", id, t.ses.ID())
	}
	switch tx.Scan(&service).Error {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsServices.create_service(): проверка логина, err=%w", err)
	}

	rg := grpc_sessions.NewRights()
	rg.SetDB(service.Rights)

	// t.ses.CherRights()
	t.data["service"] = map[string]interface{}{
		"id":      service.Id,
		"name":    service.Name,
		"key":     service.Key,
		"rights":  rg.Vue(),
		"state":   service.State == 1,
		"comment": service.Comment,
	}

	return nil
}
func (t *TSettingsServices) remove_service() error {
	var err error
	sv, ok := t.data["service"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsServices.remove_user(): не удалось привести к map")
	}

	id, ok := sv["id"].(float64)
	if !ok {
		return fmt.Errorf("TSettingsServices.remove_user(): id не установлен")
	}

	tx := db.DB.Delete(&db.Services{}, id)
	if !t.ses.Rights([]int{sessions.CAdministrator}) {
		tx.Where("user_id", t.ses.ID())
	}

	if tx.Error != nil {
		return fmt.Errorf("TSettingsServices.remove_user(): удаление, err=%w", err)
	}

	return nil
}
func (t *TSettingsServices) list() error {
	start, ok := t.data["start"].(float64)
	if !ok {
		start = 0
	}

	step, ok := t.data["step"].(float64)
	if !ok {
		step = 10
	}

	var err error
	var rows []db.Services
	tx := db.DB.Table("services").Select("id, name, comment").Offset(int(start)).Limit(int(step))
	if !t.ses.Rights([]int{sessions.CAdministrator}) {
		tx.Where("user_id = ?", t.ses.ID())
	}
	switch tx.Scan(&rows).Error {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsServices.list(): чтение таблицы, err=%w", err)
	}

	t.data["items"] = rows
	return nil
}

func (t *TSettingsServices) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TSettingsServices) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("SettingsServices", newSettingsServicesSocket)
	if err != nil {
		log.Printf("TSettingsServices(): не удалось добавить в конструктор")
	}
}
