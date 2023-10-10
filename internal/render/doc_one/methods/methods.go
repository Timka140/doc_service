package methods

import (
	"fmt"

	"projects/doc/doc_service/internal/render/doc_one/auth"
	"projects/doc/doc_service/internal/render/doc_one/methods/render"
)

type IMethods interface {
	render.IRenderDocOne
}
type TMethods struct {
	render.IRenderDocOne

	auth auth.IAuth
}

func NewDocOneMethods() (IMethods, error) {
	auth, err := auth.NewAuth()
	if err != nil {
		return nil, fmt.Errorf("NewMethods(): авторизация в doc_one, err=%w", err)
	}

	t := &TMethods{
		auth: auth,
	}

	t.IRenderDocOne = render.NewRenderDocOne(render.TRenderIN{
		Auth: auth,
	})

	return t, nil
}
