package interaction

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (t *TDocxInteraction) Connect() error {
	var err error

	rabbitHost := os.Getenv("RabbitHost")
	if rabbitHost == "" {
		return fmt.Errorf("TDocxInteraction.Connect() RabbitHost неуказан")
	}

	rabbitAuth := os.Getenv("RabbitAuth")
	if rabbitHost == "" {
		return fmt.Errorf("TXlsxInteraction.Connect() RabbitAuth неуказан")
	}

	rabbitPort := os.Getenv("RabbitPort")
	if rabbitPort == "" {
		return fmt.Errorf("TDocxInteraction.Connect() RabbitPort неуказан")
	}

	// t.conn, err = amqp.Dial("amqp://doc_service:doc_123@192.168.0.43:5672/")
	t.conn, err = amqp.Dial(fmt.Sprintf("amqp://%v@%v:%v/", rabbitAuth, rabbitHost, rabbitPort))
	if err != nil {
		return fmt.Errorf("TDocxInteraction.Connect() установка соединения, err=%w", err)
	}

	t.ch, err = t.conn.Channel()
	if err != nil {
		return fmt.Errorf("TDocxInteraction.Connect() открытие канала, err=%w", err)
	}

	return nil
}

func (t *TDocxInteraction) ConnectClose() error {
	err := t.ch.Close()
	if err != nil {
		return fmt.Errorf("TDocxInteraction.ConnectClose() закрытие канала, err=%w", err)
	}

	err = t.conn.Close()
	if err != nil {
		return fmt.Errorf("TDocxInteraction.ConnectClose() закрытие соединения, err=%w", err)
	}
	return nil
}
