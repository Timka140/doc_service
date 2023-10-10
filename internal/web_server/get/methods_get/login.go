package methods_get

import (
	"log"
	"net/http"
	"projects/doc/doc_service/internal/web_server/templates"
	"text/template"

	"github.com/gin-gonic/gin"
)

type TLoginPage struct {
	tmp *template.Template
}

func newLoginPage(in *TInGetPage) IGetPage {
	t := &TLoginPage{
		tmp: in.Tmp,
	}

	return t
}

func (t *TLoginPage) GetPath() string {
	return "/gui/login"
}

func (t *TLoginPage) GetContext(c *gin.Context) {
	var err error
	t.tmp, err = templates.NewTemplates()
	if err != nil {
		log.Printf("TLoginGet.GetContext(): инициализация шаблонов, err=%v", err)
	}

	err = t.tmp.ExecuteTemplate(c.Writer, "login.html", map[string]interface{}{})
	if err != nil {
		log.Println(err)
		c.Writer.WriteHeader(http.StatusInternalServerError)
	}
	c.Writer.WriteHeader(http.StatusOK)
}

func init() {
	err := constructors.Add("LoginPage", newLoginPage)
	if err != nil {
		log.Printf("LoginPage(): не удалось добавить в конструктор")
	}
}
