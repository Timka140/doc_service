package fill_xlsx

import (
	"bytes"
	"fmt"
	"os"

	"projects/doc/doc_service/pkg/transport/methods"

	"github.com/xuri/excelize/v2"
)

type IFillXlsx interface {
	RenderXlsx(report *methods.TReport) (err error) // Создает файл по шаблону xlsx
	ReadBytes() (data []byte, err error)            //Возвращает файл docx в байтовом виде
	WriteToPath(path string) (err error)            // Записывает файл docx по нужному пути
	GetDocument() (file methods.TFile)              // Возвращает структуру запакованного документа документа
}

type TFillXlsx struct {
	fXlsx *excelize.File

	fileXlsx *bytes.Buffer
	params   methods.TParams

	templates_path string
}

func NewFillXlsx() (IFillXlsx, error) {
	templates := os.Getenv("DocTemplates")
	if templates == "" {
		return nil, fmt.Errorf("NewFillXlsx(): значение DocTemplates в env не установлено")
	}

	t := &TFillXlsx{
		templates_path: templates,
		fileXlsx:       bytes.NewBuffer(nil),
	}

	return t, nil
}
