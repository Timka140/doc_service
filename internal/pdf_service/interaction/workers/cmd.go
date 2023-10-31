package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"projects/doc/doc_service/pkg/types"
	"time"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type TCmd struct {
	online bool

	q  amqp091.Queue
	ch *amqp.Channel
}

// initCmd() - инициализация канала
func (t *TWorker) initCmd() (types.ICmd, error) {
	cmd := &TCmd{
		ch: t.ch,
	}
	var err error

	//Декларирую канал отправки команд сервису
	cmd.q, err = t.ch.QueueDeclare(
		fmt.Sprintf("command_%v_in", t.pid), // name
		false,                               // durable
		false,                               // delete when unused
		false,                               // exclusive
		false,                               // no-wait
		nil,                                 // arguments
	)

	if err != nil {
		return nil, fmt.Errorf("TWorker.initCmd(): инициализация канала команд, err=%w", err)
	}
	return cmd, nil
}

// send() - отправка пакета данных
func (t *TCmd) send(data *[]byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := t.ch.PublishWithContext(ctx,
		"",       // exchange
		t.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        *data, //bytes,
			// AppId: uuid.NewString(),
		})
	if err != nil {
		return fmt.Errorf("TCmd.send(): публикация в очередь, err=%w", err)
	}
	return nil
}

// Exit() - завершает работу внешнего микросервиса
func (t *TCmd) Exit() error {
	d := map[string]interface{}{
		"command": "exit",
	}

	data, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("TCmd.Exit(): упаковка пакета, err=%w", err)
	}

	err = t.send(&data)
	if err != nil {
		return fmt.Errorf("TCmd.Exit(): отправка пакета, err=%w", err)
	}

	return nil
}

// Online() - статус соединения
func (t *TCmd) Online() bool {
	return true
	// return t.online
}
