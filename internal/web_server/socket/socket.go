package socket

import (
	"fmt"

	"projects/doc/doc_service/internal/web_server/sessions"
)

type TSocketValue struct {
	Data map[string]interface{}
	Ses  sessions.ISession
}

type ISocket interface {
	Start() error
	Stop() error
	Response() (map[string]interface{}, error)
	GetPid() string
}

type IMethodsSocket interface {
	Add(key string, fn func(in *TSocketValue) (ISocket, error)) error
	Get(key string) (func(in *TSocketValue) (ISocket, error), bool)
	Range(in *TSocketValue, fn func(page ISocket)) error
}

type TMethodsSocket struct {
	constructors map[string]func(in *TSocketValue) (ISocket, error)
}

var (
	constructors = &TMethodsSocket{
		constructors: make(map[string]func(in *TSocketValue) (ISocket, error)),
	}
)

func NewMethodsSocket() (IMethodsSocket, error) {
	t := constructors

	return t, nil
}
func (t *TMethodsSocket) Add(key string, fn func(in *TSocketValue) (ISocket, error)) error {
	if key == "" {
		return fmt.Errorf("TConstructorXlsx.Add(): ключ не задан")
	}

	if fn == nil {
		return fmt.Errorf("TConstructorXlsx.Add(): конструктор не задан")
	}

	t.constructors[key] = fn
	return nil
}

func (t *TMethodsSocket) Get(key string) (func(in *TSocketValue) (ISocket, error), bool) {
	if key == "" {
		return nil, false
	}

	fn, ok := t.constructors[key]
	if !ok {
		return nil, false
	}
	return fn, true
}

func (t *TMethodsSocket) Range(in *TSocketValue, fn func(page ISocket)) error {
	if fn == nil {
		return fmt.Errorf("TConstructorXlsx.Get(): ключ не задан")
	}

	for _, pConstruct := range t.constructors {
		f, err := pConstruct(in)
		if err != nil {
			return fmt.Errorf("TConstructorXlsx.Get(): создание процесса, err=%w", err)
		}
		fn(f)
	}

	return nil
}
