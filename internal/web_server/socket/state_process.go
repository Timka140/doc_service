package socket

import (
	"log"

	"github.com/google/uuid"
)

type TStateProcess struct {
	data map[string]interface{}
	pid  string
}

func newStateProcessSocket(in *TSocketValue) (ISocket, error) {
	t := &TStateProcess{
		data: in.Data,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TStateProcess) Start() error {
	for key, v := range t.data {
		log.Println(key, v)
	}
	return nil
}

func (t *TStateProcess) GetPid() string {
	return t.pid
}

func (t *TStateProcess) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TStateProcess) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("StateProcess", newStateProcessSocket)
	if err != nil {
		log.Printf("StateProcess(): не удалось добавить в конструктор")
	}
}
