package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/services"
)

type TServicesList struct {
	data map[string]interface{}
	pid  string
}

func newServicesListSocket(in *TSocketValue) (ISocket, error) {
	t := &TServicesList{
		data: in.Data,
	}

	return t, nil
}

func (t *TServicesList) Start() error {

	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TServicesList.Start(): не прочитано исполнение")
	}

	sr := services.ServicesList()

	switch execution {
	case "init":
		t.data["services"] = sr
	}

	return nil
}

func (t *TServicesList) GetPid() string {
	return t.pid
}

func (t *TServicesList) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TServicesList) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("ServicesList", newServicesListSocket)
	if err != nil {
		log.Printf("ServicesList(): не удалось добавить в конструктор")
	}
}
