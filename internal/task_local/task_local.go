package task_local

import (
	"fmt"

	"gorm.io/gorm"
)

type ITaskLocal interface {
	Init(in *TTaskLocalInit) error
}
type TTaskLocal struct {
	db        *gorm.DB
	taskID    string
	nameStore string
}

type TTaskLocalIn struct {
}

type TTaskLocalInit struct {
	TaskID string
}

func NewTaskLocal(in *TTaskLocalIn) (ITaskLocal, error) {
	t := &TTaskLocal{}

	return t, nil
}

func (t *TTaskLocal) Init(in *TTaskLocalInit) error {
	if in.TaskID == "" {
		return fmt.Errorf("TTaskLocal.Init(): TaskID не задан")
	}
	t.taskID = in.TaskID

	err := t.loadBase()
	if err != nil {
		return fmt.Errorf("TTaskLocal.Init(): создание базы, err=%w", err)
	}
	return nil
}
