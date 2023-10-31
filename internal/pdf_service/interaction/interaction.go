package interaction

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/pdf_service/interaction/workers"
	"projects/doc/doc_service/pkg/types"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IPdfInteraction interface {
	Workers() (types.IWorkers, error) //Предоставляет доступ к микросервисам
	FillPdf() (IFillPdf, error)       //Предоставляет функционал заполнения шаблонов с помощью микросервисов
	FlowPdf() (IFlowPdf, error)       //Предоставляет функционал заполнения шаблонов с помощью микросервисов
	InfoWorkers() error
}
type TPdfInteraction struct {
	sPid string
	types.IWorkers

	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewPdfInteraction() (IPdfInteraction, error) {
	t := &TPdfInteraction{
		sPid: uuid.NewString(),
	}

	err := t.Connect()
	if err != nil {
		return nil, fmt.Errorf("NewPdfInteraction(): установка соединения, err=%w", err)
	}

	t.IWorkers, err = workers.NewWorkers(&workers.TWorkersIn{
		Ch: t.ch,
	})
	if err != nil {
		return nil, fmt.Errorf("NewPdfInteraction(): инициализация микросервисов, err=%w", err)
	}

	err = t.listenInfoServers()
	if err != nil {
		return nil, fmt.Errorf("NewPdfInteraction(): ожидание информации от микросервисов, err=%w", err)
	}

	err = t.InfoWorkers()
	if err != nil {
		log.Println(err)
	}
	return t, nil
}

func (t *TPdfInteraction) Workers() (types.IWorkers, error) {
	if t.IWorkers == nil {
		return nil, fmt.Errorf("TPdfInteraction.Workers(): IWorkers не инициализирован")
	}
	return t.IWorkers, nil
}
