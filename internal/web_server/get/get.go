package get

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"projects/doc/doc_service/internal/web_server/get/methods_get"

	"github.com/gin-gonic/gin"
)

type IGet interface {
}
type TGet struct {
	router *gin.Engine
	tmp    *template.Template
}

type TInGet struct {
	Router *gin.Engine
	Tmp    *template.Template
}

func NewGet(in *TInGet) (IGet, error) {
	t := &TGet{
		router: in.Router,
		tmp:    in.Tmp,
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("NewGet(): критическая ошибка, err=%v", r)
		}
	}()

	//GET
	t.router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/gui/")
	})

	gets, err := methods_get.NewMethodsGet()
	if err != nil {
		return nil, fmt.Errorf("NewGet(): инициализация методов post, err=%w", err)
	}

	err = gets.Range(&methods_get.TInGetPage{Tmp: t.tmp}, func(page methods_get.IGetPage) {
		t.router.GET(page.GetPath(), page.GetContext)
	})
	if err != nil {
		return nil, fmt.Errorf("NewGet(): добавление routers post, err=%w", err)
	}
	return t, nil
}
