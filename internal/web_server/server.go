package web_server

import (
	"fmt"
	"os"
	"text/template"

	"projects/doc/doc_service/internal/web_server/get"
	"projects/doc/doc_service/internal/web_server/post"
	"projects/doc/doc_service/internal/web_server/static"
	"projects/doc/doc_service/internal/web_server/templates"

	"github.com/gin-gonic/gin"
)

type IServer interface {
}
type TServer struct {
	router *gin.Engine
	tmp    *template.Template
}

func NewServer() (IServer, error) {
	var err error

	// prod := os.Getenv("Prod")
	// if prod == "true" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }
	gin.SetMode(gin.ReleaseMode)

	t := &TServer{
		router: gin.New(),
		// Выключаем логи gin
		// router: gin.Default(),
	}

	// Логирование
	// t.router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	t.router.Use(gin.Recovery())

	serviceAdr := os.Getenv("ServiceAdr")
	if serviceAdr == "" {
		return nil, fmt.Errorf("NewServer(): адрес сервера не указан")
	}

	err = static.NewStatic(t.router)
	if err != nil {
		return nil, fmt.Errorf("NewServer(): инициализация статики, err=%w", err)
	}

	t.tmp, err = templates.NewTemplates()
	if err != nil {
		return nil, fmt.Errorf("NewServer(): инициализация шаблонов, err=%w", err)
	}

	_, err = get.NewGet(&get.TInGet{
		Router: t.router,
		Tmp:    t.tmp,
	})
	if err != nil {
		return nil, fmt.Errorf("NewServer(): инициализация get методов, err=%w", err)
	}

	_, err = post.NewPost(&post.TInPost{
		Router: t.router,
	})
	if err != nil {
		return nil, fmt.Errorf("NewServer(): инициализация post методов, err=%w", err)
	}

	err = t.router.Run(serviceAdr)
	if err != nil {
		return nil, fmt.Errorf("NewServer(): запуск сервера, err=%w", err)
	}

	return t, nil
}
