package interaction

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"projects/doc/doc_service/pkg/transport/formats/pdf"
	"sync"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type TPdfOut struct {
	Err  string `json:"err"`
	Data []byte `json:"data"`
}

type IFillPdf interface {
	Pack(in *pdf.TFiles) (*bytes.Buffer, error) //упаковка данных для отправки
	Send(data *bytes.Buffer) error              // отправка данных в микросервис
	Result(fn func(data *TPdfOut)) error        // принимает результат микросервиса
}

type TFillPdf struct {
	servicePid string
	ch         *amqp.Channel
	wg         sync.WaitGroup

	fIn  amqp.Queue //Канал отправки
	fOut amqp.Queue //Канал чтения
}

func (t *TPdfInteraction) FillPdf() (IFillPdf, error) {

	ch, err := t.conn.Channel() //Спорный момент возможно лучше сделать 1 канал а не открывать постоянно новый
	if err != nil {
		return nil, fmt.Errorf("TPdfInteraction.FillPdf() открытие канала, err=%w", err)
	}

	// var err error
	fillDocx := &TFillPdf{
		ch: ch,
		// ch:         t.ch,
		servicePid: uuid.NewString(),
	}

	err = fillDocx.declaration()
	if err != nil {
		return nil, fmt.Errorf("TPdfInteraction.FillPdf() декларация очередей, err=%w", err)
	}

	return fillDocx, nil
}

func (t *TFillPdf) declaration() error {
	var err error
	//In каналы отправки данных
	t.fIn, err = t.ch.QueueDeclare(
		"fill_pdf_in", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return fmt.Errorf("TFillPdf.declaration(): создание очереди, err=%w", err)
	}

	//Out каналы получения данных
	t.fOut, err = t.ch.QueueDeclare(
		fmt.Sprintf("fill_pdf_%v_out", t.servicePid), // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("TFillPdf.declaration(): создание очереди, err=%w", err)
	}

	return nil
}

// Send() - возвращает заполненный документ
func (t *TFillPdf) Send(data *bytes.Buffer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t.wg.Add(1)
	err := t.ch.PublishWithContext(ctx,
		"",         // exchange
		t.fIn.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data.Bytes(), //bytes,
			AppId:       t.servicePid,
		})
	if err != nil {
		return fmt.Errorf("TFillPdf.Send(): отправка в очередь, err=%w", err)
	}
	return nil
}

// Result() - возвращает заполненный документ
func (t *TFillPdf) Result(fn func(data *TPdfOut)) error {

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
		return fmt.Errorf("TFillPdf.Result(): прослушивание очереди, err=%w", err)
	}

	go func() {
		for d := range msgs {
			var data TPdfOut
			err := json.Unmarshal(d.Body, &data)
			if err != nil {
				log.Printf("TFillPdf.Result(): чтение пакета, err=%v", err)
				t.wg.Done()
				continue
			}

			fn(&data)
			t.wg.Done()
		}
	}()
	t.wg.Wait()

	err = t.ch.Close()
	if err != nil {
		return fmt.Errorf("TFillPdf.Result(): закрытие канала, err=%w", err)
	}

	return nil
}

// Pack() - Упаковывает данные для отправки
func (t *TFillPdf) Pack(in *pdf.TFiles) (*bytes.Buffer, error) {
	if in == nil {
		return nil, fmt.Errorf("TFillPdf.Pack(): объект пуст")
	}

	data, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("TFillPdf.Pack(): упаковка данных, err=%w", err)
	}

	return bytes.NewBuffer(data), nil
}
