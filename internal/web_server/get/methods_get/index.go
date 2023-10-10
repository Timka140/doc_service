package methods_get

import (
	"log"
	"net/http"
	"projects/doc/doc_service/internal/web_server/sessions"
	"projects/doc/doc_service/internal/web_server/templates"
	"text/template"

	"github.com/gin-gonic/gin"
)

type TIndexPage struct {
	// ctx *gin.Context
	tmp *template.Template
}

func newIndexPage(in *TInGetPage) IGetPage {
	t := &TIndexPage{
		tmp: in.Tmp,
	}

	return t
}

func (t *TIndexPage) GetPath() string {
	return "/gui/"
}

func (t *TIndexPage) GetContext(c *gin.Context) {
	ses := sessions.Ses.GetSes(c)
	if ses == nil {
		return
	}

	var err error
	t.tmp, err = templates.NewTemplates()
	if err != nil {
		log.Printf("NewServer(): инициализация шаблонов, err=%v", err)
	}

	page := make(map[string]interface{})

	err = t.tmp.ExecuteTemplate(c.Writer, "index.html", page)
	if err != nil {
		log.Println(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func init() {
	err := constructors.Add("IndexPage", newIndexPage)
	if err != nil {
		log.Printf("IndexPage(): не удалось добавить в конструктор")
	}
}
