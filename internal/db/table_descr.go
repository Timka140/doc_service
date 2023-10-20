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
	Id      int64
	Path    string
	Name    string
	Tp      int
	Comment string
	Data    []byte
	Hash    string
}

type Template struct {
	Id       int64
	PathBase string
	TaskID   int64
	Name     string
}
