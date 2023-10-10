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
