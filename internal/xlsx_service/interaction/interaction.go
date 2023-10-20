package interaction

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/xlsx_service/interaction/workers"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IXlsxInteraction interface {
	Workers() (workers.IWorkers, error) //Предоставляет доступ к микросервисам
	FillXlsx() (IFillXlsx, error)       //Предоставляет функционал заполнения шаблонов с помощью микросервисов
	FlowXlsx() (IFlowXlsx, error)       //Предоставляет функционал заполнения шаблонов с помощью микросервисов
	InfoWorkers() error
}
type TXlsxInteraction struct {
	sPid string
	workers.IWorkers

	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewXlsxInteraction() (IXlsxInteraction, error) {
	t := &TXlsxInteraction{
		sPid: uuid.NewString(),
	}

	err := t.Connect()
	if err != nil {
		return nil, fmt.Errorf("NewXlsxInteraction(): установка соединения, err=%w", err)
	}

	t.IWorkers, err = workers.NewWorkers(&workers.TWorkersIn{
		Ch: t.ch,
	})
	if err != nil {
		return nil, fmt.Errorf("NewXlsxInteraction(): инициализация микросервисов, err=%w", err)
	}

	err = t.listenInfoServers()
	if err != nil {
		return nil, fmt.Errorf("NewXlsxInteraction(): ожидание информации от микросервисов, err=%w", err)
	}

	err = t.InfoWorkers()
	if err != nil {
		log.Println(err)
	}
	return t, nil
}

func (t *TXlsxInteraction) Workers() (workers.IWorkers, error) {
	if t.IWorkers == nil {
		return nil, fmt.Errorf("TXlsxInteraction.Workers(): IWorkers не инициализирован")
	}
	return t.IWorkers, nil
}
