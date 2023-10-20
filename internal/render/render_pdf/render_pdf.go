package render_pdf

import (
	"fmt"
	"projects/doc/doc_service/internal/render/render_pdf/fill_pdf"
)

// Комплекс модулей для обработки файлов формата pdf
type IRenderPdf interface {
	fill_pdf.IFillPdf
}

type TRenderPdf struct {
	fill_pdf.IFillPdf
}

func NewRenderPdf() (IRenderPdf, error) {
	var err error

	t := &TRenderPdf{}

	t.IFillPdf, err = fill_pdf.NewFillPdf()
	if err != nil {
		return nil, fmt.Errorf("NewRenderPdf(): инициализация NewFillPdf, err=%w", err)
	}

	return t, nil
}
