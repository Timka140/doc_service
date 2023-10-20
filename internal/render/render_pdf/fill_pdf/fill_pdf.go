package fill_pdf

import (
	"bytes"

	"projects/doc/doc_service/internal/pdf_service"
	"projects/doc/doc_service/internal/pdf_service/interaction"

	"projects/doc/doc_service/pkg/transport/methods"
)

type IFillPdf interface {
	RenderPdf(report *methods.TReport) (err error) // Создает файл по шаблону pdf
	ReadBytes() (data []byte, err error)           //Возвращает файл pdf в байтовом виде
	WriteToPath(path string) (err error)           // Записывает файл pdf по нужному пути
	GetDocument() (file methods.TFile)             // Возвращает структуру запакованного документа документа
}

type TFillPdf struct {
	filePdf *bytes.Buffer
	params  methods.TParams
	flow    interaction.IFlowPdf
}

func NewFillPdf() (IFillPdf, error) {
	t := &TFillPdf{
		filePdf: bytes.NewBuffer(nil),
		flow:    pdf_service.FlowPdf,
	}

	return t, nil
}
