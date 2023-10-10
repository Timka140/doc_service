package constructors_xlsx

import (
	"bytes"
	"fmt"

	formats_xlsx "projects/doc/doc_service/pkg/transport/formats/xlsx"

	"github.com/xuri/excelize/v2"
)

// Структура конструктора
type TDependenciesRender struct {
	FXlsx *excelize.File
}

type IConstructorsXlsx interface {
	SetData(in *formats_xlsx.TValuesRender) error
	Render() error
	GetFileXlsx() (*bytes.Buffer, error)
}

// Методы управления конвеиером
type IConstructorXlsx interface {
	Add(key string, fn func(in *TDependenciesRender) IConstructorsXlsx) error
	Get(key string) (func(in *TDependenciesRender) IConstructorsXlsx, error)
}

type TConstructorXlsx struct {
	constructors map[string]func(in *TDependenciesRender) IConstructorsXlsx
}

var (
	constructors = &TConstructorXlsx{
		constructors: make(map[string]func(in *TDependenciesRender) IConstructorsXlsx),
	}
)

func NewConstructor() IConstructorXlsx {
	t := constructors
	return t
}

func (t *TConstructorXlsx) Add(key string, fn func(in *TDependenciesRender) IConstructorsXlsx) error {
	if key == "" {
		return fmt.Errorf("TConstructorXlsx.Add(): ключ не задан")
	}

	if fn == nil {
		return fmt.Errorf("TConstructorXlsx.Add(): конструктор не задан")
	}

	t.constructors[key] = fn
	return nil
}

func (t *TConstructorXlsx) Get(key string) (func(in *TDependenciesRender) IConstructorsXlsx, error) {
	if key == "" {
		return nil, fmt.Errorf("TConstructorXlsx.Get(): ключ не задан")
	}

	fn, ok := t.constructors[key]
	if !ok {
		return nil, fmt.Errorf("TConstructorXlsx.Get(): конструктор не найден")
	}
	return fn, nil
}
