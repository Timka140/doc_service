package methods_post

import (
	"bytes"
	"os"

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

type TTemplateModified struct {
	// ctx *gin.Context
}

func newTemplateModified(in *TInPostPage) IPostPage {
	t := &TTemplateModified{}

	return t

}

func (t *TTemplateModified) GetPath() string {
	return "/template_modified"
}

func (t *TTemplateModified) GetContext(c *gin.Context) {
	ses := sessions.Ses.GetSes(c)
	if ses == nil {
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		log.Println("TTemplateModified.GetContext(): чтение формы, err=%w", err)
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
			log.Println("TTemplateModified.GetContext(): чтение файла, err=%w", err)
			return
		}
		name = file.Filename
		io.Copy(&data, f)
		f.Close()
	}

	catalog := filepath.Join("store/template", template_id)
	err = os.MkdirAll(catalog, 0755)
	if err != nil {
		log.Println("TTemplateModified.GetContext(): создание папки, err=%w", err)
	}

	pFile := filepath.Join(catalog, name)
	f, err := os.Create(pFile)
	if err != nil {
		log.Println("TTemplateModified.GetContext(): создание файла, err=%w", err)
		return
	}
	defer f.Close()
	_, err = f.Write(data.Bytes())
	if err != nil {
		log.Println("TTemplateModified.GetContext(): запись в файл, err=%w", err)
		return
	}

	tmp := &types.TFile{
		Name:         name,
		Update:       time.Now().Format(time.RFC3339Nano),
		Ext:          filepath.Ext(name),
		PathTemplate: pFile,
	}

	pack, err := template.Pack(tmp)
	if err != nil {
		log.Println("TTemplateModified.GetContext(): упаковка, err=%w", err)
	}

	err = db.DB.Table("templates").Where("id = ?", template_id).Update("data", pack).Error
	if err != nil {
		log.Println("TTemplateModified.GetContext(): не удалось загрузить файл, err=%w", err)
	}

	err = ses.SendMessage(map[string]interface{}{
		"tp": "UpdateTemplate",
	})
	if err != nil {
		log.Println("TTemplateModified.GetContext(): обновление шаблона, err=%w", err)
	}

	c.Status(http.StatusOK)
}

func init() {
	err := constructors.Add("TemplateModified", newTemplateModified)
	if err != nil {
		log.Printf("LoadTaskData(): не удалось добавить в конструктор")
	}
}
