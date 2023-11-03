package db

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// _ "github.com/mattn/go-sqlite3"
)

const (
	file_sql = "store.db3" //:locked.sqlite?cache=shared
)

var (
	DB *gorm.DB
)

type TDB struct {
	*gorm.DB
	sync.RWMutex
}

func NewDB() error {
	t := &TDB{}
	db, err := gorm.Open(sqlite.Open(file_sql), &gorm.Config{})
	// db, err := sql.Open("sqlite3", file_sql)
	if err != nil {
		return fmt.Errorf("NewDB(): открытие базы данных, err=%w", err)
	}
	t.DB = db
	DB = db.Session(&gorm.Session{})

	err = t.create_tables()
	if err != nil {
		return fmt.Errorf("NewDB(): добавление таблиц, err=%w", err)
	}

	err = t.default_user()
	if err != nil {
		return fmt.Errorf("NewDB(): добавление пользователя, err=%w", err)
	}

	// err = t.filter_fields_default()
	// if err != nil {
	// 	return fmt.Errorf("NewDB(): загрузка переменных по умолчанию, err=%w", err)
	// }

	return nil
}

func (db *TDB) create_tables() error {
	rows, err := db.Table("sqlite_master").Select("name").Where("type=?", "table").Rows()
	// rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table';`)
	if err != nil {
		return fmt.Errorf("create_tables(): проверка таблиц в базе, err=%w", err)
	}

	tables := make(map[string]bool)
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return fmt.Errorf("create_tables(): чтение таблиц с базы, err=%w", err)
		}

		tables[name] = true
	}

	tables_create := map[string]string{
		"users": `CREATE TABLE "users" (
			"id"	INTEGER NOT NULL UNIQUE,
			"login"	TEXT NOT NULL,
			"password"	BLOB NOT NULL,
			"rights"	BLOB NOT NULL,
			"state"	INTEGER NOT NULL DEFAULT 0,
			PRIMARY KEY("id" AUTOINCREMENT)
		);`,
		"tasks": `CREATE TABLE "tasks" (
			"id"	INTEGER NOT NULL UNIQUE,
			"path"	TEXT,
			"tp"	INTEGER,
			"name"	TEXT,
			"comment"	TEXT,
			"base"	BLOB,
			"tp_task"	INTEGER,
			PRIMARY KEY("id" AUTOINCREMENT)
		);`,
		"task": `CREATE TABLE "task" (
			"id"	INTEGER NOT NULL UNIQUE,
			"path_base"	TEXT,
			"name"	TEXT,
			"task_id"	INTEGER NOT NULL,
			PRIMARY KEY("id" AUTOINCREMENT)
		);`,
		"templates": `CREATE TABLE "templates" (
			"id"	INTEGER NOT NULL UNIQUE,
			"path"	TEXT,
			"name"	TEXT,
			"tp"	INTEGER,
			"comment"	TEXT,
			"data"	BLOB,
			"hash"	TEXT,
			PRIMARY KEY("id" AUTOINCREMENT)
		);`,
		"template_versions": `CREATE TABLE "template_versions" (
			"id"	INTEGER NOT NULL UNIQUE,
			"data"	BLOB,
			"template_id"	INTEGER NOT NULL,
			"user"	TEXT,
			"date_update"	TEXT,
			PRIMARY KEY("id" AUTOINCREMENT)
		);`,
	}

	for key, val := range tables_create {
		_, ok := tables[key]
		if ok {
			continue
		}

		db.RLock()
		if err = db.Exec(val, nil).Error; err != nil {
			return fmt.Errorf("create_tables(): добавление таблицы пользователей, err=%w", err)
		}
		db.RUnlock()
		// DB.RLock()
		// if _, err := db.Exec(val).Error; err != nil {
		// 	return fmt.Errorf("create_tables(): добавление таблицы пользователей, err=%w", err)
		// }
		// DB.RUnlock()
	}

	return nil
}
func (db *TDB) default_user() error {
	var err error
	// if _, err = os.Stat(file_sql); !errors.Is(err, os.ErrNotExist) {
	// 	return nil
	// }

	var id sql.NullInt64
	err = db.Table("users").Select("id").Where("login = ?", "bondarenkotg").Scan(&id).Error
	// err = db.QueryRow("SELECT id FROM users WHERE login = 'bondarenkotg'").Scan(&id)
	switch err {
	case sql.ErrNoRows:
	case nil:
		return nil
	default:
		return fmt.Errorf("default_user(): поиск пользователя по умолчанию, err=%w", err)
	}

	login := "bondarenkotg"
	password := "B~G|sUGP7bN%"
	key := md5.Sum([]byte(fmt.Sprintf("%v:docGenerator:%v", login, password)))
	hash := hex.EncodeToString(key[:])

	err = db.Table("users").Create(Users{
		Login:    "bondarenkotg",
		Password: hash,
		State:    1,
		Rights:   nil,
	}).Error
	if err != nil {
		return fmt.Errorf("default_user(): запись пользователя в базу, err=%w", err)
	}
	// DB.RLock()
	// _, err = DB.Exec(`INSERT INTO users (login, password, state, rights) VALUES (?, ?, 1, '{}')`, "bondarenkotg", hash)
	// if err != nil {
	// 	return fmt.Errorf("default_user(): запись пользователя в базу, err=%w", err)
	// }
	// DB.RUnlock()
	return nil
}

// func (db *TDB) filter_fields_default() error {
// 	var err error
// 	var id sql.NullInt64
// 	err = db.Table("filter_fields").Select("id_processes").Where("id_processes = ?", 0).Scan(&id).Error
// 	// err = db.DB.QueryRow("SELECT id_processes FROM filter_fields WHERE id_processes = 0").Scan(&id)
// 	switch err {
// 	case sql.ErrNoRows:
// 	case nil:
// 		return nil
// 	default:
// 		return fmt.Errorf("filter_fields_default(): поиск пользователя по умолчанию, err=%w", err)
// 	}

// 	filterFields := map[string]interface{}{
// 		"fields": map[string]interface{}{
// 			"organization_name": map[string]interface{}{
// 				"default_value": "______________________________",
// 				"required":      true,
// 				"position":      1,
// 				"forced":        false,
// 			},
// 			"organization_full_name": map[string]interface{}{
// 				"default_value": "____________________________________________________________",
// 				"required":      true,
// 				"position":      1,
// 				"forced":        false,
// 			},
// 			"organization_fio": map[string]interface{}{
// 				"default_value": "______________________________",
// 				"required":      true,
// 				"position":      1,
// 				"forced":        false,
// 			},
// 			"organization_document": map[string]interface{}{
// 				"default_value": "_____________________",
// 				"required":      true,
// 				"position":      1,
// 				"forced":        false,
// 			},
// 			"organization_post": map[string]interface{}{
// 				"default_value": "______________________________",
// 				"required":      true,
// 				"position":      1,
// 				"forced":        false,
// 			},
// 			"organization_initials": map[string]interface{}{
// 				"default_value": "______________________________",
// 				"required":      false,
// 				"position":      1,
// 				"forced":        true,
// 			},
// 			"estimated_voltage_level": map[string]interface{}{
// 				"default_value": "В соответствии с законода-тельством РФ",
// 				"required":      false,
// 				"position":      1,
// 				"forced":        true,
// 			},
// 		},
// 	}

// 	pack, err := json.Marshal(filterFields)
// 	if err != nil {
// 		return fmt.Errorf("filter_fields_default(): запись значений по умолчанию, err=%w", err)
// 	}

// 	DB.RLock()
// 	_, err = db.DB.Exec(`INSERT INTO filter_fields (id_processes, data) VALUES (?, ?)`, 0, pack)
// 	if err != nil {
// 		return fmt.Errorf("filter_fields_default(): запись пользователя в базу, err=%w", err)
// 	}
// 	DB.RUnlock()
// 	return nil
// }
