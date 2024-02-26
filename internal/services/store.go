package services

import (
	"fmt"
	"projects/doc/doc_service/pkg/transport/connect"
)

// Структура конструктора
type TInServices struct {
	Sid    string
	Create *connect.TCreate
}

type IService interface {
	Info() map[string]interface{}
	SetInfo(*connect.TInfo)

	Ping() int64
	SetPing(ping int64)

	Data() map[string]interface{}

	Name() string
	Comment() string
}

type IServices interface {
	Add(key string, srv IService) error
	Get(key string) (IService, error)
	Range(fn func(srv IService)) error
	Delete(key string)
}

type TServices struct {
	constructors map[string]IService
}

var (
	constructors = &TServices{
		constructors: make(map[string]IService),
	}
)

func New() (IServices, error) {
	t := constructors

	return t, nil
}
func (t *TServices) Add(key string, srv IService) error {
	if key == "" {
		return fmt.Errorf("TServices.Add(): ключ не задан")
	}

	if srv == nil {
		return fmt.Errorf("TServices.Add(): конструктор не задан")
	}

	t.constructors[key] = srv
	return nil
}

func (t *TServices) Get(key string) (IService, error) {
	if key == "" {
		return nil, fmt.Errorf("TServices.Get(): ключ не задан")
	}

	fn, ok := t.constructors[key]
	if !ok {
		return nil, fmt.Errorf("TServices.Get(): конструктор не найден, constructor = %v", key)
	}
	return fn, nil
}

func (t *TServices) Delete(key string) {
	delete(t.constructors, key)
}

func (t *TServices) Range(fn func(srv IService)) error {
	if fn == nil {
		return fmt.Errorf("TServices.Get(): ключ не задан")
	}

	for _, pConstruct := range t.constructors {
		if pConstruct == nil {
			continue
		}
		fn(pConstruct)
	}

	return nil
}
