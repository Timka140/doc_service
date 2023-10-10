package socket

import (
	"log"

	"github.com/google/uuid"
)

type TTaskLocal struct {
	data map[string]interface{}
	pid  string
}

func newTaskLocalSocket(in *TSocketValue) (ISocket, error) {
	t := &TTaskLocal{
		data: in.Data,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TTaskLocal) Start() error {
	for key, v := range t.data {
		log.Println(key, v)
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
