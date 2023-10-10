package interaction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type TDocxIn struct {
	Template []byte                 `json:"template"`
	Data     map[string]interface{} `json:"data"`
}

type TDocxOut struct {
	Err  *string `json:"err"`
	Data []byte  `json:"data"`
}

type IFillDocx interface {
	Pack(in *TDocxIn) (*bytes.Buffer, error) //упаковка данных для отправки
	Send(data *bytes.Buffer) error           // отправка данных в микросервис
	Result(fn func(data *TDocxOut)) error    // принимает результат микросервиса
}

type TFillDocx struct {
	servicePid string
	ch         *amqp.Channel
	wg         sync.WaitGroup

	fIn  amqp.Queue //Канал отправки
	fOut amqp.Queue //Канал чтения
}

func (t *TDocxInteraction) FillDocx() (IFillDocx, error) {

	ch, err := t.conn.Channel() //Спорный момент возможно лучше сделать 1 канал а не открывать постоянно новый
	if err != nil {
		return nil, fmt.Errorf("TDocxInteraction.FillDocx() открытие канала, err=%w", err)
	}

	fillDocx := &TFillDocx{
		ch:         ch,
		servicePid: uuid.NewString(),
	}

	err = fillDocx.declaration()
	if err != nil {
		return nil, fmt.Errorf("TDocxInteraction.FillDocx() декларация очередей, err=%w", err)
	}

	return fillDocx, nil
}

func (t *TFillDocx) declaration() error {
	var err error
	//In каналы отправки данных
	t.fIn, err = t.ch.QueueDeclare(
		"fill_docx_in", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return fmt.Errorf("TDocxInteraction.declaration(): создание очереди, err=%w", err)
	}

	//Out каналы получения данных
	t.fOut, err = t.ch.QueueDeclare(
		fmt.Sprintf("fill_docx_%v_out", t.servicePid), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("TDocxInteraction.declaration(): создание очереди, err=%w", err)
	}

	return nil
}

// Send() - возвращает заполненный документ
func (t *TFillDocx) Send(data *bytes.Buffer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.wg.Add(1)
	err := t.ch.PublishWithContext(ctx,
		"",         // exchange
		t.fIn.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			Body:          data.Bytes(), //bytes,
			CorrelationId: t.servicePid,
		})
	if err != nil {
		return fmt.Errorf("TDocxInteraction.FillDocx(): отправка в очередь, err=%w", err)
	}
	return nil
}

// Result() - возвращает заполненный документ
func (t *TFillDocx) Result(fn func(data *TDocxOut)) error {

	msgs, err := t.ch.Consume(
		t.fOut.Name, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return fmt.Errorf("TFillDocx.Result(): прослушивание очереди, err=%w", err)
	}

	go func() {
		for d := range msgs {
			var data TDocxOut
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				log.Printf("TFillDocx.Result(): чтение пакета, err=%v", err)
				t.wg.Done()
				continue
			}

			fn(&data)
			t.wg.Done()
		}
	}()
	t.wg.Wait()

	err = t.ch.Cancel("", true)
	if err != nil {
		return fmt.Errorf("TFillDocx.Result(): прерывание канала, err=%w", err)
	}

	// err = t.ch.Close()
	// if err != nil {
	// 	return fmt.Errorf("TFillDocx.Result(): закрытие канала, err=%w", err)
	// }

	return nil
}

// Pack() - Упаковывает данные для отправки
func (t *TFillDocx) Pack(in *TDocxIn) (*bytes.Buffer, error) {
	if in == nil {
		return nil, fmt.Errorf("TFillDocx.Pack(): объект пуст")
	}

	data, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("TFillDocx.Pack(): упаковка данных, err=%w", err)
	}

	return bytes.NewBuffer(data), nil
}
