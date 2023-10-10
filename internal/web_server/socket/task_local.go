package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/task_local"

	"github.com/google/uuid"
)

type TTaskLocal struct {
	data map[string]interface{}
	pid  string
	task task_local.ITaskLocal
}

func newTaskLocalSocket(in *TSocketValue) (ISocket, error) {
	task, err := task_local.NewTaskLocal(&task_local.TTaskLocalIn{})
	if err != nil {
		return nil, fmt.Errorf("newTaskLocalSocket(): инициализация task_local, err=%w", err)
	}

	t := &TTaskLocal{
		data: in.Data,
		pid:  uuid.NewString(),
		task: task,
	}

	return t, nil
}

func (t *TTaskLocal) Start() error {
	var err error
	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TTaskLocal.Start(): не прочитано исполнение")
	}

	task_id, ok := t.data["task_id"].(string)
	if !ok {
		return fmt.Errorf("TTaskLocal.Start(): task_id не задан")
	}

	switch execution {
	case "init":
		err = t.task.Init(&task_local.TTaskLocalInit{
			TaskID: task_id,
		})
		if err != nil {
			return fmt.Errorf("TTaskLocal.Start(): инициализация task, err=%w", err)
		}
	}

	return nil
}

func (t *TTaskLocal) GetPid() string {
	return t.pid
}

func (t *TTaskLocal) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TTaskLocal) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("TaskLocal", newTaskLocalSocket)
	if err != nil {
		log.Printf("TaskLocal(): не удалось добавить в конструктор")
	}
}
