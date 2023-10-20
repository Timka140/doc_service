package workers

import (
	"fmt"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var workers sync.Map //список подключенных приложений

type IWorkers interface {
	Add(pid string) error           //Добавляет микросервис
	List() TListWorkers             //Возвращает список микросервисов
	Get(pid string) (IWorker, bool) //Получить микросервис
	Range(fn func(pid string, work IWorker))
	Delete(pid string)
}
type TWorkers struct {
	ch *amqp.Channel
}

type TWorkersIn struct {
	Ch *amqp.Channel
}

func NewWorkers(in *TWorkersIn) (IWorkers, error) {
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

type TListWorker struct {
	Pid    string
	Online bool
}
type TListWorkers map[string]*TListWorker

func (t *TWorkers) List() TListWorkers {
	list := make(TListWorkers)

	workers.Range(func(key, value any) bool {
		pid, ok := key.(string)
		if !ok {
			return true
		}

		work, ok := value.(IWorker)
		if !ok {
			return true
		}

		list[pid] = &TListWorker{
			Pid:    pid,
			Online: work.Online(),
		}
		return true
	})
	return list
}

func (t *TWorkers) Get(pid string) (IWorker, bool) {
	v, ok := workers.Load(pid)
	if !ok {
		return nil, false
	}

	work, ok := v.(IWorker)
	if !ok {
		return nil, false
	}

	return work, true
}

func (t *TWorkers) Delete(pid string) {
	workers.Delete(pid)
}

func (t *TWorkers) Range(fn func(pid string, work IWorker)) {

	workers.Range(func(key, value any) bool {
		pid, ok := key.(string)
		if !ok {
			return true
		}

		work, ok := value.(IWorker)
		if !ok {
			return true
		}

		fn(pid, work)

		return true
	})
}
