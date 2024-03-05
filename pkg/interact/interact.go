package interact

import "fmt"

type IInteract interface {
	Login(login, password string) error
}

type TInteract struct {
	Protocol string
	Address  string

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
