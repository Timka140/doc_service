package interaction

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/docx_service/interaction/workers"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IDocxInteraction interface {
	Workers() (workers.IWorkers, error) //Предоставляет доступ к микросервисам
	FillDocx() (IFillDocx, error)       //Предоставляет функционал заполнения шаблонов с помощью микросервисов
	InfoWorkers() error
}
type TDocxInteraction struct {
	sPid string
	workers.IWorkers

	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewDocxInteraction() (IDocxInteraction, error) {
	t := &TDocxInteraction{
		sPid: uuid.NewString(),
	}

	err := t.Connect()
	if err != nil {
		return nil, fmt.Errorf("NewDocxInteraction(): установка соединения, err=%w", err)
	}

	t.IWorkers, err = workers.NewWorkers(&workers.TWorkersIn{
		Ch: t.ch,
	})
	if err != nil {
		return nil, fmt.Errorf("NewDocxInteraction(): инициализация микросервисов, err=%w", err)
	}

	err = t.listenInfoServers()
	if err != nil {
		return nil, fmt.Errorf("NewDocxInteraction(): ожидание информации от микросервисов, err=%w", err)
	}

	err = t.InfoWorkers()
	if err != nil {
		log.Println(err)
	}
	return t, nil
}

func (t *TDocxInteraction) Workers() (workers.IWorkers, error) {
	if t.IWorkers == nil {
		return nil, fmt.Errorf("Workers(): IWorkers не инициализирован")
	}
	return t.IWorkers, nil
}
