package xlsx_service

import (
	"fmt"
	"log"
	"os"
	"projects/doc/doc_service/internal/xlsx_service/interaction"
	"projects/doc/doc_service/internal/xlsx_service/service"
	"sync"

	"github.com/google/uuid"
)

var XlsxServices IXlsxService
var FlowXlsx interaction.IFlowXlsx

type IXlsxService interface {
	interaction.IXlsxInteraction
	StartService() error //Запуск микросервиса
	StopServices()       //Остановка микросервисов
}
type TDocxService struct {
	interaction.IXlsxInteraction
	services   sync.Map
	rabbitHost string
	rabbitPort string
}

func NewXlsxService() (IXlsxService, error) {
	rabbitHost := os.Getenv("RabbitHost")
	if rabbitHost == "" {
		return nil, fmt.Errorf("NewXlsxService() RabbitHost неуказан")
	}

	rabbitPort := os.Getenv("RabbitPort")
	if rabbitPort == "" {
		return nil, fmt.Errorf("NewXlsxService() RabbitPort неуказан")
	}

	t := &TDocxService{
		rabbitHost: rabbitHost,
		rabbitPort: rabbitPort,
	}

	var err error
	t.IXlsxInteraction, err = interaction.NewXlsxInteraction()
	if err != nil {
		return nil, fmt.Errorf("NewXlsxService() инициализация модулю взаимодействия, err=%w", err)
	}

	return t, nil
}

// StartService() - добавляет сервис для обработки шаблонов
func (t *TDocxService) StartService() error {
	pid := uuid.NewString()
	s, err := service.NewService(&service.TInStart{
		RabbitHost: t.rabbitHost,
		RabbitPort: t.rabbitPort,
		Pid:        pid,
	})
	if err != nil {
		return fmt.Errorf("StartService(): создание микросервиса, err=%w", err)
	}

	err = s.Start()
	if err != nil {
		return fmt.Errorf("StartService(): запуск микросервиса, err=%w", err)
	}

	t.services.Store(pid, s)

	return nil

}

// StopServices() - остановка всех сервисов
func (t *TDocxService) StopServices() {
	t.services.Range(func(key, value any) bool {
		s, ok := value.(service.IService)
		if !ok {
			return true
		}
		err := s.Stop()
		if err != nil {
			log.Printf("StopServices(): остановка сервиса, pid: %v, err:%v", key, err)
		}
		return true
	})
}

func init() {
	var err error
	XlsxServices, err = NewXlsxService()
	if err != nil {
		log.Printf("docx_service.init(): инициализация сервисов, err:%v", err)
		os.Exit(1)
	}

	FlowXlsx, err = XlsxServices.FlowXlsx()
	if err != nil {
		log.Printf("docx_service.init(): инициализация FlowXlsx, err=%v", err)
		os.Exit(1)
	}

}
