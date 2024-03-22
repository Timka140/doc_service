package socket

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/google/uuid"
)

type TSettingsUsers struct {
	data map[string]interface{}
	ses  sessions.ISession
	pid  string
}

func newSettingsUsersSocket(in *TSocketValue) (ISocket, error) {
	t := &TSettingsUsers{
		data: in.Data,
		ses:  in.Ses,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TSettingsUsers) Start() error {
	var err error
	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TSettingsUsers.Start(): не прочитано исполнение")
	}

	switch execution {
	case "check_login":
		err = t.check_login()
		if err != nil {
			return fmt.Errorf("TSettingsUsers.Start(): проверка логина, err=%w", err)
		}
	case "create_user":
		err = t.create_user()
		if err != nil {
			return fmt.Errorf("TSettingsUsers.Start(): создание пользователя, err=%w", err)
		}
	case "update_user":
		err = t.update_user()
		if err != nil {
			return fmt.Errorf("TSettingsUsers.Start(): обновление пользователя, err=%w", err)
		}
	case "info_user":
		err = t.info_user()
		if err != nil {
			return fmt.Errorf("TSettingsUsers.Start(): получение данных пользователя, err=%w", err)
		}
	case "remove_user":
		err = t.remove_user()
		if err != nil {
			return fmt.Errorf("TSettingsUsers.Start(): удаление пользователя, err=%w", err)
		}
	case "list":
		err = t.list()
		if err != nil {
			return fmt.Errorf("TSettingsUsers.Start(): список пользователей, err=%w", err)
		}
	}
	return nil
}

func (t *TSettingsUsers) GetPid() string {
	return t.pid
}

func (t *TSettingsUsers) check_login() error {
	if !t.ses.Rights([]int{sessions.CAdministrator}) {
		return nil
	}

	login, ok := t.data["login"].(string)
	if !ok {
		return fmt.Errorf("TSettingsUsers.check_login(): приведение к строке, login")
	}

	var err error
	var row db.Users
	err = db.DB.Table("users").Select("id").Where("login = ?", login).Scan(&row).Error
	switch err {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsUsers.check_login(): проверка логина, err=%w", err)
	}
	t.data["present_login"] = 2
	if row.Id != 0 {
		t.data["present_login"] = 1
	}

	return nil
}
func (t *TSettingsUsers) create_user() error {
	if !t.ses.Rights([]int{sessions.CAdministrator}) {
		return nil
	}
	var err error
	var user db.Users

	us, ok := t.data["user"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsUsers.create_user(): не удалось привести к map, err=%w", err)
	}

	user.Login, ok = us["login"].(string)
	if !ok || user.Login == "" {
		return fmt.Errorf("TSettingsUsers.create_user(): некорректный логин, err=%w", err)
	}

	password, ok := us["password"].(string)
	if !ok {
		return fmt.Errorf("TSettingsUsers.create_user(): некорректный пароль, err=%w", err)
	}

	key := md5.Sum([]byte(fmt.Sprintf("%v:docGenerator:%v", user.Login, password)))
	user.Password = hex.EncodeToString(key[:])
	user.State = 1

	rights, ok := us["rights"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsUsers.create_user(): некорректный пароль, err=%w", err)
	}

	rg := sessions.NewRights()

	r, err := rg.Set(rights)
	if err != nil {
		return fmt.Errorf("TSettingsUsers.create_user(): установка прав, err=%w", err)
	}
	user.Rights = r.Get()

	user.Comment, ok = us["comment"].(string)
	if !ok {
		return fmt.Errorf("TSettingsUsers.create_user(): некорректный пароль, err=%w", err)
	}

	err = db.DB.Table("users").Create(&user).Error
	switch err {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsUsers.create_user(): создание пользователя, err=%w", err)
	}
	return nil
}
func (t *TSettingsUsers) update_user() error {
	if !t.ses.Rights([]int{sessions.CAdministrator}) {
		return nil
	}

	var err error
	var user db.Users

	us, ok := t.data["user"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsUsers.update_user(): не удалось привести к map, err=%w", err)
	}

	id, ok := us["id"].(float64)
	if !ok {
		return fmt.Errorf("TSettingsUsers.update_user(): id не установлен, err=%w", err)
	}

	user.Login, ok = us["login"].(string)
	if !ok || user.Login == "" {
		return fmt.Errorf("TSettingsUsers.update_user(): некорректный логин, err=%w", err)
	}

	password, ok := us["password"].(string)
	if ok {
		key := md5.Sum([]byte(fmt.Sprintf("%v:docGenerator:%v", user.Login, password)))
		user.Password = hex.EncodeToString(key[:])
	}

	rights, ok := us["rights"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsUsers.update_user(): некорректный пароль, err=%w", err)
	}

	rg := sessions.NewRights()

	r, err := rg.Set(rights)
	if err != nil {
		return fmt.Errorf("TSettingsUsers.update_user(): установка прав, err=%w", err)
	}
	user.Rights = r.Get()

	user.Comment, ok = us["comment"].(string)
	if !ok {
		return fmt.Errorf("TSettingsUsers.update_user(): некорректный пароль, err=%w", err)
	}

	err = db.DB.Table("users").Where("id = ?", id).Updates(user).Error
	switch err {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsUsers.create_user(): обновление пользователя, err=%w", err)
	}

	return nil
}

func (t *TSettingsUsers) info_user() error {
	if !t.ses.Rights([]int{sessions.CAdministrator}) {
		return nil
	}
	var err error
	us, ok := t.data["user"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsUsers.update_user(): не удалось привести к map, err=%w", err)
	}

	id, ok := us["id"].(float64)
	if !ok {
		return fmt.Errorf("TSettingsUsers.update_user(): id не установлен, err=%w", err)
	}

	var user db.Users
	err = db.DB.Table("users").Select("id, login, rights, comment").Where("id = ?", id).Scan(&user).Error
	switch err {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsUsers.create_user(): получение информации пользователя, err=%w", err)
	}

	rg := sessions.NewRights()
	rg.SetDB(user.Rights)

	// t.ses.CherRights()
	t.data["user"] = map[string]interface{}{
		"id":      user.Id,
		"login":   user.Login,
		"rights":  rg.Vue(),
		"comment": user.Comment,
	}

	return nil
}
func (t *TSettingsUsers) remove_user() error {
	if !t.ses.Rights([]int{sessions.CAdministrator}) {
		return nil
	}
	var err error
	us, ok := t.data["user"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TSettingsUsers.remove_user(): не удалось привести к map")
	}

	id, ok := us["id"].(float64)
	if !ok {
		return fmt.Errorf("TSettingsUsers.remove_user(): id не установлен")
	}

	err = db.DB.Delete(&db.Users{}, id).Error
	if !ok {
		return fmt.Errorf("TSettingsUsers.remove_user(): удаление, err=%w", err)
	}

	return nil
}
func (t *TSettingsUsers) list() error {
	start, ok := t.data["start"].(float64)
	if !ok {
		start = 0
	}

	step, ok := t.data["step"].(float64)
	if !ok {
		step = 10
	}

	var err error
	var rows []db.Users
	err = db.DB.Table("users").Select("id, login, comment").Offset(int(start)).Limit(int(step)).Scan(&rows).Error
	switch err {
	case nil:
	case sql.ErrNoRows:
	default:
		return fmt.Errorf("TSettingsUsers.list(): чтение таблицы, err=%w", err)
	}

	t.data["items"] = rows
	return nil
}

func (t *TSettingsUsers) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TSettingsUsers) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("SettingsUsers", newSettingsUsersSocket)
	if err != nil {
		log.Printf("TSettingsUsers(): не удалось добавить в конструктор")
	}
}
