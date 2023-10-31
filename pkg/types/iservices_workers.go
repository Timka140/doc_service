package types

type TListWorker struct {
	Pid    string
	Online bool
}

type TListWorkers map[string]*TListWorker

type IWorker interface {
	ICmd
}

type ICmd interface {
	Exit() error // завершает работу внешнего микросервиса
	Online() bool
}

type IWorkers interface {
	Add(pid string) error           //Добавляет микросервис
	List() TListWorkers             //Возвращает список микросервисов
	Len() int                       //Возвращает длину
	Get(pid string) (IWorker, bool) //Получить микросервис
	Range(fn func(pid string, work IWorker))
	Delete(pid string)
}
