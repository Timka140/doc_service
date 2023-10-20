package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/pdf_service"
	"strconv"
	"sync"

	"github.com/google/uuid"
)

type TRunPdfServices struct {
	data map[string]interface{}
	pid  string
}

func newRunPdfServicesSocket(in *TSocketValue) (ISocket, error) {
	t := &TRunPdfServices{
		data: in.Data,
		pid:  uuid.NewString(),
	}

	return t, nil
}

func (t *TRunPdfServices) Start() error {
	var err error
	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TListDocxServices.Start(): не прочитано исполнение")
	}

	switch execution {
	case "start":
		err = t.startServices()
		if err != nil {
			return fmt.Errorf("TRunDocxServices.Start(): запуск сервисов, err=%w", err)
		}
	case "stop":
		err = t.closeServices()
		if err != nil {
			return fmt.Errorf("TRunDocxServices.Start(): остановка микросервисов, err=%w", err)
		}
	}

	return nil
}

func (t *TRunPdfServices) GetPid() string {
	return t.pid
}

func (t *TRunPdfServices) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TRunPdfServices) Stop() error {

	return nil
}

func (t *TRunPdfServices) closeServices() error {
	works, err := pdf_service.PdfServices.Workers()
	if err != nil {
		return fmt.Errorf("TRunPdfServices.closeAllServices(): получение микросервисов, err=%w", err)
	}

	services, ok := t.data["services"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TRunPdfServices.closeAllServices(): получение списка микросервисов, err=%w", err)
	}

	for pid, val := range services {
		work, ok := works.Get(pid)
		if !ok {
			continue
		}
		v, ok := val.(map[string]interface{})
		if !ok {
			continue
		}

		s, ok := v["select"].(bool)
		if !ok {
			continue
		}
		if !s {
			continue
		}

		err := work.Exit()
		if err != nil {
			return fmt.Errorf("TRunPdfServices.closeAllServices(): получение микросервисов, err=%v", err)
		}
		works.Delete(pid)
	}

	err = pdf_service.PdfServices.InfoWorkers()
	if err != nil {
		return fmt.Errorf("TRunPdfServices.closeAllServices(): запрос количества сервисов, %w", err)
	}

	return nil
}

func (t *TRunPdfServices) startServices() error {
	qs, ok := t.data["quantity"].(string)
	if !ok {
		return fmt.Errorf("TRunPdfServices.Start(): quantity не строка, %v", t.data["quantity"])
	}

	quantity, err := strconv.Atoi(qs)
	if err != nil {
		return fmt.Errorf("TRunPdfServices.Start(): quantity не число, %w", err)
	}
	var wg sync.WaitGroup
	for i := 0; i < quantity; i++ {
		wg.Add(1)
		go func() {
			err = pdf_service.PdfServices.StartService()
			if err != nil {
				log.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	err = pdf_service.PdfServices.InfoWorkers()
	if err != nil {
		return fmt.Errorf("TRunPdfServices.startServices(): запрос количества сервисов, %w", err)
	}

	return nil
}

func init() {
	err := constructors.Add("RunPdfServices", newRunPdfServicesSocket)
	if err != nil {
		log.Printf("RunPdfServices(): не удалось добавить в конструктор")
	}
}
