package pdf_service

import (
	"fmt"
	"log"
	"os"
	"projects/doc/doc_service/internal/pdf_service/interaction"
	"projects/doc/doc_service/internal/pdf_service/service"
	"sync"

	"github.com/google/uuid"
)

var PdfServices IPdfService
var FlowPdf interaction.IFlowPdf

type IPdfService interface {
	interaction.IPdfInteraction
	StartService() error //Запуск микросервиса
	StopServices()       //Остановка микросервисов
}
type TPdfService struct {
	interaction.IPdfInteraction
	services   sync.Map
	rabbitHost string
	rabbitPort string
}

func NewPdfService() (IPdfService, error) {
	rabbitHost := os.Getenv("RabbitHost")
	if rabbitHost == "" {
		return nil, fmt.Errorf("NewPdfService() RabbitHost неуказан")
	}

	rabbitPort := os.Getenv("RabbitPort")
	if rabbitPort == "" {
		return nil, fmt.Errorf("NewPdfService() RabbitPort неуказан")
	}

	t := &TPdfService{
		rabbitHost: rabbitHost,
		rabbitPort: rabbitPort,
	}

	var err error
	t.IPdfInteraction, err = interaction.NewPdfInteraction()
	if err != nil {
		return nil, fmt.Errorf("NewPdfService() инициализация модулю взаимодействия, err=%w", err)
	}

	return t, nil
}

// StartService() - добавляет сервис для обработки шаблонов
func (t *TPdfService) StartService() error {
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
func (t *TPdfService) StopServices() {
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
	PdfServices, err = NewPdfService()
	if err != nil {
		log.Printf("pdf_service.init(): инициализация сервисов, err:%v", err)
		os.Exit(1)
	}

	FlowPdf, err = PdfServices.FlowPdf()
	if err != nil {
		log.Printf("pdf_service.init(): инициализация FlowPdf, err=%v", err)
		os.Exit(1)
	}

}
