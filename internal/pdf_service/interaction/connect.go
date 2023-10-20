package interaction

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (t *TPdfInteraction) Connect() error {
	var err error

	rabbitHost := os.Getenv("RabbitHost")
	if rabbitHost == "" {
		return fmt.Errorf("TXlsxInteraction.Connect() RabbitHost неуказан")
	}

	rabbitPort := os.Getenv("RabbitPort")
	if rabbitPort == "" {
		return fmt.Errorf("TXlsxInteraction.Connect() RabbitPort неуказан")
	}

	t.conn, err = amqp.Dial(fmt.Sprintf("amqp://guest:guest@%v:%v/", rabbitHost, rabbitPort))
	if err != nil {
		return fmt.Errorf("TXlsxInteraction.Connect() установка соединения, err=%w", err)
	}

	t.ch, err = t.conn.Channel()
	if err != nil {
		return fmt.Errorf("TXlsxInteraction.Connect() открытие канала, err=%w", err)
	}

	return nil
}

func (t *TPdfInteraction) ConnectClose() error {
	err := t.ch.Close()
	if err != nil {
		return fmt.Errorf("TXlsxInteraction.ConnectClose() закрытие канала, err=%w", err)
	}

	err = t.conn.Close()
	if err != nil {
		return fmt.Errorf("TXlsxInteraction.ConnectClose() закрытие соединения, err=%w", err)
	}
	return nil
}
