package interaction

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IFlowPdf interface {
	Send(pack []byte) (*TPdfOut, error) // отправка данных в микросервис
	Close() error
}
type tWorkData struct {
	Data *[]byte
}
type TWork struct {
	Pid  string
	data chan tWorkData //Ожидает ответа выполнения
}
type TFlowPdf struct {
	servicePid string
	ch         *amqp.Channel
	wg         sync.WaitGroup
	works      sync.Map

	fIn  amqp.Queue //Канал отправки
	fOut amqp.Queue //Канал чтения
}

func (t *TPdfInteraction) FlowPdf() (IFlowPdf, error) {

	ch, err := t.conn.Channel() //Спорный момент возможно лучше сделать 1 канал а не открывать постоянно новый
	if err != nil {
		return nil, fmt.Errorf("TPdfInteraction.TFlowPdf() открытие канала, err=%w", err)
	}

	// var err error
	fillDocx := &TFlowPdf{
		ch: ch,
		// ch:         t.ch,
		servicePid: uuid.NewString(),
	}

	err = fillDocx.declaration()
	if err != nil {
		return nil, fmt.Errorf("TPdfInteraction.TFlowPdf() декларация очередей, err=%w", err)
	}

	err = fillDocx.result()
	if err != nil {
		return nil, fmt.Errorf("TPdfInteraction.TFlowPdf() слушатель, err=%w", err)
	}

	return fillDocx, nil
}

func (t *TFlowPdf) declaration() error {
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
		return fmt.Errorf("TFlowPdf.declaration(): создание очереди, err=%w", err)
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
		return fmt.Errorf("TFlowPdf.declaration(): создание очереди, err=%w", err)
	}

	return nil
}

// Send() - возвращает заполненный документ
func (t *TFlowPdf) Send(pack []byte) (*TPdfOut, error) {

	if pack == nil {
		return nil, fmt.Errorf("TFlowPdf.Send(): объект пуст")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pid := uuid.NewString()
	work := &TWork{
		Pid:  pid,
		data: make(chan tWorkData),
	}

	t.works.Store(pid, work)
	defer func() {
		close(work.data) //
		t.works.Delete(pid)
	}()

	t.wg.Add(1)
	err := t.ch.PublishWithContext(ctx,
		"",         // exchange
		t.fIn.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        pack, //bytes,
			MessageId:   pid,
			AppId:       t.servicePid,
		})
	if err != nil {
		return nil, fmt.Errorf("TFlowPdf.Send(): отправка в очередь, err=%w", err)
	}

	wData := <-work.data
	if wData.Data == nil {
		return nil, fmt.Errorf("TFlowPdf.Send(): ошибка чтения результата")
	}

	file := &TPdfOut{}
	err = json.Unmarshal(*wData.Data, file)
	if err != nil {
		return nil, fmt.Errorf("TFlowPdf.Send(): чтение пакета, err=%v", err)
	}

	return file, nil
}

// result() - возвращает заполненный документ
func (t *TFlowPdf) result() error {

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
		return fmt.Errorf("TFlowPdf.result(): прослушивание очереди, err=%w", err)
	}

	go func() {
		for d := range msgs {
			if d.MessageId == "" {
				log.Printf("TFlowPdf.result(): сообщение без id")
				continue
			}
			data, ok := t.works.Load(d.MessageId)
			if !ok {
				log.Printf("TFlowPdf.result(): данные не найдены")
				continue
			}

			work, ok := data.(*TWork)
			if !ok {
				log.Printf("TFlowPdf.result(): неизвестный тип данных")
				continue
			}

			work.data <- tWorkData{
				Data: &d.Body,
			}
			t.wg.Done()
		}
	}()

	return nil
}

// Close() - Закрытие потока
func (t *TFlowPdf) Close() error {
	t.wg.Wait() // Ожидаем завершения всех задач
	err := t.ch.Close()
	if err != nil {
		return fmt.Errorf("TFlowPdf.Result(): закрытие канала, err=%w", err)
	}

	return err
}
