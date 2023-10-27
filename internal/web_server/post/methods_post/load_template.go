package methods_post

import (
	"bytes"

	"io"
	"log"
	"net/http"
	"path/filepath"
	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/internal/template"
	"projects/doc/doc_service/internal/web_server/sessions"
	"projects/doc/doc_service/pkg/types"
	"time"

	"github.com/gin-gonic/gin"
)

type TLoadTemplateData struct {
	// ctx *gin.Context
}

func newLoadTemplateData(in *TInPostPage) IPostPage {
	t := &TLoadTemplateData{}

	return t

}

func (t *TLoadTemplateData) GetPath() string {
	return "/load_template"
}

func (t *TLoadTemplateData) GetContext(c *gin.Context) {
	ses := sessions.Ses.GetSes(c)
	if ses == nil {
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		log.Println("TLoadTemplateData.GetContext(): чтение формы, err=%w", err)
		return
	}

	form, _ := c.MultipartForm()

	template_id := form.Value["template_id"][0]
	file := form.File["file"]

	if template_id == "" {
		return
	}

	if len(file) != 1 {
		return
	}

	var name string
	var data bytes.Buffer

	for _, file := range file {
		f, err := file.Open()
		if err != nil {
			log.Println("TLoadTemplateData.GetContext(): чтение файла, err=%w", err)
			return
		}
		name = file.Filename
		io.Copy(&data, f)
		f.Close()
	}

	tmp := &types.TFile{
		Name:   name,
		Data:   data.Bytes(),
		Update: time.Now().Format(time.RFC3339Nano),
		Ext:    filepath.Ext(name),
	}

	pack, err := template.Pack(tmp)
	if err != nil {
		log.Println("TLoadTemplateData.GetContext(): упаковка, err=%w", err)
	}

	err = db.DB.Table("templates").Where("id = ?", template_id).Update("data", pack).Error
	if err != nil {
		log.Println("TLoadTemplateData.GetContext(): не удалось загрузить файл, err=%w", err)
	}

	c.Status(http.StatusOK)
}

func init() {
	err := constructors.Add("LoadTemplateData", newLoadTemplateData)
	if err != nil {
		log.Printf("LoadTaskData(): не удалось добавить в конструктор")
	}
}
