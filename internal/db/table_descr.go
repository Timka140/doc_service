package db

import (
	"github.com/lib/pq"
)

// Таблица пользователь
type Users struct {
	Id       int64
	Login    string
	Password string
	State    int64         `gorm:"type:int"`
	Rights   pq.Int64Array `gorm:"type:[]int64"`
	Comment  string
}

type Services struct {
	Id      int64
	Name    string
	Key     string        `gorm:"type:string"`
	UserID  int64         `gorm:"type:int64;column:user_id"`
	State   int64         `gorm:"type:int"`
	Rights  pq.Int64Array `gorm:"type:[]int64"`
	Comment string
}

type Tasks struct {
	Id      int64
	Path    string
	Tp      int
	Name    string
	Comment string
}

type Task struct {
	Id       int64
	PathBase string
	TaskID   int64
	Name     string
}

type Templates struct {
	Id           int64
	Path         string
	Name         string
	Tp           int
	Comment      string
	Data         []byte
	Hash         string
	UserID       int64 `gorm:"type:int64;column:user_id"`
	Organization int64 `gorm:"type:int64;column:organization"`
}

type TemplateVersions struct {
	Id         int64
	Data       []byte
	TemplateID int64
	User       string
	DateUpdate string
}
