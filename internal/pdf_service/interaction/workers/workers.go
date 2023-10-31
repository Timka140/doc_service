package workers

import (
	"fmt"
	"projects/doc/doc_service/pkg/types"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var workers sync.Map //список подключенных приложений

type TWorkers struct {
	ch *amqp.Channel
}

type TWorkersIn struct {
	Ch *amqp.Channel
}

func NewWorkers(in *TWorkersIn) (types.IWorkers, error) {
	t := &TWorkers{
		ch: in.Ch,
	}

	return t, nil
}

func (t *TWorkers) Add(pid string) error {
	work, err := t.newWorker(&TWorkerIn{
		Pid: pid,
	})
	if err != nil {
		return fmt.Errorf("TWorkers.Add(): инициализация работы, err=%w", err)
	}
	workers.Store(pid, work)
	return nil
}

func (t *TWorkers) List() types.TListWorkers {
	list := make(types.TListWorkers)

	workers.Range(func(key, value any) bool {
		pid, ok := key.(string)
		if !ok {
			return true
		}

		work, ok := value.(types.IWorker)
		if !ok {
			return true
		}

		list[pid] = &types.TListWorker{
			Pid:    pid,
			Online: work.Online(),
		}
		return true
	})
	return list
}

func (t *TWorkers) Len() int {
	var len int
	workers.Range(func(key, value any) bool {
		len++
		return true
	})
	return len
}

func (t *TWorkers) Get(pid string) (types.IWorker, bool) {
	v, ok := workers.Load(pid)
	if !ok {
		return nil, false
	}

	work, ok := v.(types.IWorker)
	if !ok {
		return nil, false
	}

	return work, true
}

func (t *TWorkers) Delete(pid string) {
	workers.Delete(pid)
}

func (t *TWorkers) Range(fn func(pid string, work types.IWorker)) {

	workers.Range(func(key, value any) bool {
		pid, ok := key.(string)
		if !ok {
			return true
		}

		work, ok := value.(types.IWorker)
		if !ok {
			return true
		}

		fn(pid, work)

		return true
	})
}
