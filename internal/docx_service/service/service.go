package service

import (
	"fmt"
	"os/exec"
	"sync"
)

type TClose struct {
	Close bool
}

type IService interface {
	Start() error
	Stop() error
}

type TService struct {
	cmd *exec.Cmd
	wg  sync.WaitGroup

	cClose chan TClose

	err error

	host string
	port string
	auth string
	pid  string
}
type TInStart struct {
	RabbitHost string
	RabbitPort string
	RabbitAuth string
	Pid        string
}

func NewService(in *TInStart) (IService, error) {
	t := &TService{
		host:   in.RabbitHost,
		port:   in.RabbitPort,
		auth:   in.RabbitAuth,
		pid:    in.Pid,
		cClose: make(chan TClose, 1),
	}

	return t, nil
}

func (t *TService) Start() error {
	err := t.startService()
	if err != nil {
		return fmt.Errorf("NewStream(): запуск сервиса, err=%w", err)
	}

	go t.listen()
	return nil
}

func (t *TService) Stop() error {
	if t.cClose == nil {
		return fmt.Errorf("TStart.Stop(): канал закрыт")
	}

	t.cClose <- TClose{
		Close: true,
	}

	t.wg.Wait()
	if t.err != nil {
		return t.err
	}
	return nil
}
