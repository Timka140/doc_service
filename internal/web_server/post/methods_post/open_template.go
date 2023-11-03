package methods_post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"projects/doc/doc_service/internal/template"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/gin-gonic/gin"
)

type TOpenTemplate struct {
	// ctx *gin.Context
}

func newOpenTemplate(in *TInPostPage) IPostPage {
	t := &TOpenTemplate{}

	return t
}

func (t *TOpenTemplate) GetPath() string {
	return "/open_template"
}

func (t *TOpenTemplate) GetContext(c *gin.Context) {
	ses := sessions.Ses.GetSes(c)
	if ses == nil {
		return
	}

	data := bytes.NewBuffer(nil)
	io.Copy(data, c.Request.Body)
	if data.Len() == 0 {
		return
	}
	params := make(map[string]interface{})
	err := json.Unmarshal(data.Bytes(), &params)
	// err := c.Request.ParseForm()
	if err != nil {
		log.Printf("TLoginPost.GetContext(): чтение параметров, err=%v", err)
		return
	}

	template_id := params["template_id"]

	tmp, err := template.New(template_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": fmt.Sprintf("TOpenTemplate.GetContext(): инициализация шаблона, err=%v", err),
		})

	}

	if !tmp.IsFile() {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": "TOpenTemplate.GetContext(): шаблон не загружен",
		})
	}

	file, err := tmp.BaseLoad()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": fmt.Sprintf("TOpenTemplate.GetContext(): не удалось получить шаблон, err=%v", err),
		})
	}

	tData, err := tmp.Template()
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"err": "TOpenTemplate.GetContext(): чтение шаблона",
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"file": tData,
		"ext":  file.Ext,
		"name": file.Name,
	})
}

func init() {
	err := constructors.Add("OpenTemplate", newOpenTemplate)
	if err != nil {
		log.Printf("OpenTemplate(): не удалось добавить в конструктор")
	}
}
