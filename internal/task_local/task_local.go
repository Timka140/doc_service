package task_local

import "gorm.io/gorm"

type ITaskLocal interface{}
type TTaskLocal struct {
	db        *gorm.DB
	nameStore string
}

func NewTaskLocal() (ITaskLocal, error) {
	t := &TTaskLocal{}
	return t, nil
}
