package render

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"projects/doc/doc_service/internal/render/doc_one/auth"
	"projects/doc/doc_service/pkg/transport/methods"
)

type TRenderDocOne struct {
	auth auth.IAuth
	data map[string][]string

	templateID string
	format     string
	link       string

	body []byte

	params        methods.TParams
	fileDocOne    *bytes.Buffer
	fileExtDocOne string
}

type IRenderDocOne interface {
	RenderDocOne(report *methods.TReport) error
	Stop() error

	WriteToPath(path string) (err error) // Записывает файл по нужному пути
	ReadBytes() (data []byte, err error) // Возвращает файл в байтовом виде
	GetDocument() (file methods.TFile)   // Возвращает структуру запакованного документа документа

}

type TRenderIN struct {
	Auth auth.IAuth
}

func NewRenderDocOne(in TRenderIN) IRenderDocOne {
	t := &TRenderDocOne{
		auth:       in.Auth,
		fileDocOne: bytes.NewBuffer(nil),
	}
	return t
}

func (t *TRenderDocOne) render() error {

	//api/v3/documents?TemplateID=b72cedea-61b6-11ed-9e58-02c16399bff6
	fmt.Printf("%v/api/v3/documents?TemplateID=%v", t.auth.GetURL(), t.templateID)
	r, err := http.NewRequest("POST", fmt.Sprintf("%v/api/v3/documents?TemplateID=%v", t.auth.GetURL(), t.templateID), bytes.NewBuffer(t.body))
	if err != nil {
		return fmt.Errorf("Render(): формирование запроса: %w", err)
	}

	r.Header.Add("Authorization", fmt.Sprintf("Basic %v", t.auth.GetAuthBase64()))
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("Render(): отправка запроса: %w", err)
	}

	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	rData := make(map[string]interface{})
	err = json.Unmarshal(resp, &rData)
	if err != nil {
		log.Println(err)
	}

	var ok bool
	t.link, ok = rData["Link"].(string)
	if !ok {
		return fmt.Errorf("Render(): чтение ссылки: %v", rData["Link"])
	}

	t.format, ok = rData["DownloadFormats"].(string)
	if !ok {
		return fmt.Errorf("Render(): чтение формата: %v", rData["DownloadFormats"])
	}

	return nil
}
