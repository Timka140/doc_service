package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/pdf_service"

	"github.com/google/uuid"
)

type TListPdfServices struct {
	data map[string]interface{}
	pid  string
}

func newListPdfServicesSocket(in *TSocketValue) (ISocket, error) {
	t := &TListPdfServices{
		data: in.Data,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TListPdfServices) Start() error {

	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TListPdfServices.Start(): не прочитано исполнение")
	}

	workers, err := pdf_service.PdfServices.Workers()
	if err != nil {
		log.Println(err)
	}

	switch execution {
	case "list":
		t.data["services"] = workers.List()
	case "info":
		err := pdf_service.PdfServices.InfoWorkers()
		if err != nil {
			log.Println(err)
		}
		t.data["services"] = workers.List()
	}

	return nil
}

func (t *TListPdfServices) GetPid() string {
	return t.pid
}

func (t *TListPdfServices) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TListPdfServices) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("ListPdfServices", newListPdfServicesSocket)
	if err != nil {
		log.Printf("ListPdfServices(): не удалось добавить в конструктор")
	}
}
