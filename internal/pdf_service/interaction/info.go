package interaction

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"projects/doc/doc_service/internal/web_server/sessions"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// InfoWorkers() - получает информацию о уже запущенных микросервисах
func (t *TPdfInteraction) InfoWorkers() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := t.ch.ExchangeDeclare(
		"pdf_info_in", // name
		"fanout",      // type
		false,         // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return fmt.Errorf("TPdfInteraction.InfoServers(): создание пространства для обмена, err=%w", err)
	}

	err = t.ch.PublishWithContext(ctx,
		"pdf_info_in", // exchange
		"",            // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte{}, //bytes,
			// AppId: uuid.NewString(),
		})
	if err != nil {
		return fmt.Errorf("TPdfInteraction.InfoServers(): публикация в очередь, err=%w", err)
	}

	return nil
}

// listenInfoServers() - ожидает информацию от сервисов
func (t *TPdfInteraction) listenInfoServers() error {

	err := t.ch.ExchangeDeclare(
		"pdf_info_out", // name
		"fanout",       // type
		false,          // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return fmt.Errorf("TPdfInteraction.listenInfoServers(): создание пространства для обмена, err=%w", err)
	}

	q, err := t.ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("TPdfInteraction.listenInfoServers(): создание очереди, err=%w", err)
	}

	err = t.ch.QueueBind(
		q.Name,         // queue name
		"",             // routing key
		"pdf_info_out", // exchange
		false,
		nil)
	if err != nil {
		return fmt.Errorf("TPdfInteraction.listenInfoServers(): подключение очереди к пространству, err=%w", err)
	}

	msgs, err := t.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("TPdfInteraction.listenInfoServers(): прослушивание очереди, err=%w", err)
	}

	// var wg sync.WaitGroup
	// wg.Add(1)

	type tInfo struct {
		Pid    string `json:"pid"`
		Online bool   `json:"online"`
	}
	go func() {
		for d := range msgs {
			var info tInfo
			err := json.Unmarshal(d.Body, &info)
			if err != nil {
				log.Println("TPdfInteraction.listenInfoServers(): чтение пакета, err=%w", err)
				continue
			}

			if info.Online {
				err = t.IWorkers.Add(info.Pid)
				if err != nil {
					log.Println("TPdfInteraction.listenInfoServers(): добавление микросервиса, err=%w", err)
				}
			} else {
				t.IWorkers.Delete(info.Pid)
			}

			//Рассылка списка сервисов
			sessions.Ses.RangeSes(func(ses sessions.ISession) {
				if ses.GetCurrentPage() != "/gui/services/pdf" {
					return
				}
				ses.SendMessage(map[string]interface{}{
					"tp":       "ListPdfServices",
					"services": t.IWorkers.List(),
				})
			})
			// log.Printf("Received a message: %s %s", d.CorrelationId, d.Body)
		}
	}()
	// wg.Wait()

	return nil
}
