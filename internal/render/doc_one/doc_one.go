package doc_one

import (
	"fmt"

	"projects/doc/doc_service/internal/render/doc_one/methods"
)

type IDocOne interface {
	methods.IMethods
}
type TDocOne struct {
	methods.IMethods
}

func NewDocOne() (IDocOne, error) {
	var err error
	t := &TDocOne{}

	t.IMethods, err = methods.NewDocOneMethods()
	if err != nil {
		return nil, fmt.Errorf("NewDocOne(): установка методов, err=%w", err)
	}

	return t, nil
}
