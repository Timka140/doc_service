package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/docx_service"

	"github.com/google/uuid"
)

type TListDocxServices struct {
	data map[string]interface{}
	pid  string
}

func newListDocxServicesSocket(in *TSocketValue) (ISocket, error) {
	t := &TListDocxServices{
		data: in.Data,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TListDocxServices) Start() error {

	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TListDocxServices.Start(): не прочитано исполнение")
	}
	workers, err := docx_service.DocxServices.Workers()
	if err != nil {
		log.Println(err)
	}

	switch execution {
	case "list":
		t.data["services"] = workers.List()
	case "info":
		err := docx_service.DocxServices.InfoWorkers()
		if err != nil {
			log.Println(err)
		}
		t.data["services"] = workers.List()
	}

	return nil
}

func (t *TListDocxServices) GetPid() string {
	return t.pid
}

func (t *TListDocxServices) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TListDocxServices) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("ListDocxServices", newListDocxServicesSocket)
	if err != nil {
		log.Printf("ListDocxServices(): не удалось добавить в конструктор")
	}
}
