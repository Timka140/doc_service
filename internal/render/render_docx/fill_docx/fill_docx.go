package fill_docx

import (
	"bytes"
	"fmt"
	"os"

	"projects/doc/doc_service/pkg/transport/methods"
)

type IFillDocx interface {
	RenderDocx(report *methods.TReport) (err error) // Создает файл по шаблону docx
	ReadBytes() (data []byte, err error)            //Возвращает файл docx в байтовом виде
	WriteToPath(path string) (err error)            // Записывает файл docx по нужному пути
	GetDocument() (file methods.TFile)              // Возвращает структуру запакованного документа документа
}

type TFillDocx struct {
	fileDocx *bytes.Buffer
	params   methods.TParams

	templates_path string
}

func NewFillDocx() (IFillDocx, error) {
	templates := os.Getenv("DocTemplates")
	if templates == "" {
		return nil, fmt.Errorf("NewFillDocx(): значение DocTemplates в env не установлено")
	}

	t := &TFillDocx{
		templates_path: templates,
		fileDocx:       bytes.NewBuffer(nil),
	}

	return t, nil
}
