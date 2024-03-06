package interact

import (
	"bytes"
	"fmt"
)

type IInteract interface {
	Login(login, password string) error
	Close()

	LoadTemplate(id, name string, file *bytes.Buffer) error
}

type TInteract struct {
	Protocol string
	Address  string

	close chan bool

	token string
}

func New(protocol, address string) (IInteract, error) {
	if address == "" {
		return nil, fmt.Errorf("interact.New(): адрес не указан")
	}

	if protocol == "" {
		protocol = "http"
	}

	t := &TInteract{
		Protocol: protocol,
		Address:  address,
	}

	return t, nil
}
