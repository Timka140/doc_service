package render_docx

import (
	"fmt"

	"projects/doc/doc_service/internal/render/render_docx/fill_docx"
)

// Комплекс модулей для обработки файлов формата docx
type IRenderDocx interface {
	fill_docx.IFillDocx
}

type TRenderDocx struct {
	fill_docx.IFillDocx
}

func NewRenderDocx() (IRenderDocx, error) {
	var err error
	t := &TRenderDocx{}

	t.IFillDocx, err = fill_docx.NewFillDocx()
	if err != nil {
		return nil, fmt.Errorf("NewRenderDocx(): инициализация NewFillDocx, err=%w", err)
	}

	return t, nil
}
