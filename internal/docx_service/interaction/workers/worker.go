package workers

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type IWorker interface {
	ICmd
}
type TWorker struct {
	ICmd
	pid string
	ch  *amqp.Channel
}
type TWorkerIn struct {
	Pid string
}

func (t *TWorkers) newWorker(in *TWorkerIn) (IWorker, error) {
	work := &TWorker{
		pid: in.Pid,
		ch:  t.ch,
	}

	var err error
	work.ICmd, err = work.initCmd()
	if err != nil {
		return nil, fmt.Errorf("NewWorkers(): инициализация команд, err=%w", err)
	}

	return work, nil
}
