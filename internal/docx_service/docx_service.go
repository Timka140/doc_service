package docx_service

import (
	"fmt"
	"log"
	"os"
	"projects/doc/doc_service/internal/docx_service/interaction"
	"projects/doc/doc_service/internal/docx_service/service"
	"sync"

	"github.com/google/uuid"
)

var DocxServices IDocxService
var FlowDocx interaction.IFlowDocx

type IDocxService interface {
	interaction.IDocxInteraction
	StartService() error //Запуск микросервиса
	StopServices()       //Остановка микросервисов
}
type TDocxService struct {
	interaction.IDocxInteraction
	services   sync.Map
	rabbitHost string
	rabbitAuth string
	rabbitPort string
}

func NewDocxService() (IDocxService, error) {
	rabbitHost := os.Getenv("RabbitHost")
	if rabbitHost == "" {
		return nil, fmt.Errorf("NewDocxService() RabbitHost неуказан")
	}

	rabbitAuth := os.Getenv("RabbitAuth")
	if rabbitHost == "" {
		return nil, fmt.Errorf("NewDocxService() RabbitAuth неуказан")
	}

	rabbitPort := os.Getenv("RabbitPort")
	if rabbitPort == "" {
		return nil, fmt.Errorf("NewDocxService() RabbitPort неуказан")
	}

	t := &TDocxService{
		rabbitHost: rabbitHost,
		rabbitAuth: rabbitAuth,
		rabbitPort: rabbitPort,
	}

	var err error
	t.IDocxInteraction, err = interaction.NewDocxInteraction()
	if err != nil {
		return nil, fmt.Errorf("NewDocxService() инициализация модулю взаимодействия, err=%w", err)
	}

	return t, nil
}

// StartService() - добавляет сервис для обработки шаблонов
func (t *TDocxService) StartService() error {
	pid := uuid.NewString()
	s, err := service.NewService(&service.TInStart{
		RabbitHost: t.rabbitHost,
		RabbitPort: t.rabbitPort,
		RabbitAuth: t.rabbitAuth,
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
	DocxServices, err = NewDocxService()
	if err != nil {
		log.Printf("docx_service.init(): инициализация сервисов, err:%v", err)
		os.Exit(1)
	}

	FlowDocx, err = DocxServices.FlowDocx()
	if err != nil {
		log.Printf("docx_service.init(): инициализация FlowDocx, err=%v", err)
		os.Exit(1)
	}

}
