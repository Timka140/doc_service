package methods_post

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Структура конструктора
type TInPostPage struct {
	*gin.Context
}

type IPostPage interface {
	GetPath() string
	GetContext(ctx *gin.Context)
}

type IMethodsPost interface {
	Add(key string, fn func(in *TInPostPage) IPostPage) error
	Get(key string) (func(in *TInPostPage) IPostPage, error)
	Range(in *TInPostPage, fn func(page IPostPage)) error
}

type TMethodsPost struct {
	constructors map[string]func(in *TInPostPage) IPostPage
}

var (
	constructors = &TMethodsPost{
		constructors: make(map[string]func(in *TInPostPage) IPostPage),
	}
)

func NewMethodsPost() (IMethodsPost, error) {
	t := constructors

	return t, nil
}
func (t *TMethodsPost) Add(key string, fn func(in *TInPostPage) IPostPage) error {
	if key == "" {
		return fmt.Errorf("TConstructorXlsx.Add(): ключ не задан")
	}

	if fn == nil {
		return fmt.Errorf("TConstructorXlsx.Add(): конструктор не задан")
	}

	t.constructors[key] = fn
	return nil
}

func (t *TMethodsPost) Get(key string) (func(in *TInPostPage) IPostPage, error) {
	if key == "" {
		return nil, fmt.Errorf("TConstructorXlsx.Get(): ключ не задан")
	}

	fn, ok := t.constructors[key]
	if !ok {
		return nil, fmt.Errorf("TConstructorXlsx.Get(): конструктор не найден")
	}
	return fn, nil
}

func (t *TMethodsPost) Range(in *TInPostPage, fn func(page IPostPage)) error {
	if fn == nil {
		return fmt.Errorf("TConstructorXlsx.Get(): ключ не задан")
	}

	for _, pConstruct := range t.constructors {
		p := pConstruct(in)
		fn(p)
	}

	return nil
}
