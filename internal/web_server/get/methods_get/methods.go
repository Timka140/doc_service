package methods_get

import (
	"fmt"
	"text/template"

	"github.com/gin-gonic/gin"
)

// Структура конструктора
type TInGetPage struct {
	// *gin.Context
	Tmp *template.Template
}

type IGetPage interface {
	GetPath() string
	GetContext(ctx *gin.Context)
}

type IMethodsGet interface {
	Add(key string, fn func(in *TInGetPage) IGetPage) error
	Get(key string) (func(in *TInGetPage) IGetPage, error)
	Range(in *TInGetPage, fn func(page IGetPage)) error
}

type TMethodsGet struct {
	constructors map[string]func(in *TInGetPage) IGetPage
}

var (
	constructors = &TMethodsGet{
		constructors: make(map[string]func(in *TInGetPage) IGetPage),
	}
)

func NewMethodsGet() (IMethodsGet, error) {
	t := constructors

	return t, nil
}
func (t *TMethodsGet) Add(key string, fn func(in *TInGetPage) IGetPage) error {
	if key == "" {
		return fmt.Errorf("TConstructorXlsx.Add(): ключ не задан")
	}

	if fn == nil {
		return fmt.Errorf("TConstructorXlsx.Add(): конструктор не задан")
	}

	t.constructors[key] = fn
	return nil
}

func (t *TMethodsGet) Get(key string) (func(in *TInGetPage) IGetPage, error) {
	if key == "" {
		return nil, fmt.Errorf("TConstructorXlsx.Get(): ключ не задан")
	}

	fn, ok := t.constructors[key]
	if !ok {
		return nil, fmt.Errorf("TConstructorXlsx.Get(): конструктор не найден")
	}
	return fn, nil
}

func (t *TMethodsGet) Range(in *TInGetPage, fn func(page IGetPage)) error {
	if fn == nil {
		return fmt.Errorf("TConstructorXlsx.Get(): ключ не задан")
	}

	for _, pConstruct := range t.constructors {
		p := pConstruct(in)
		fn(p)
	}

	return nil
}
