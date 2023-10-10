package render_xlsx

import (
	"fmt"

	"projects/doc/doc_service/internal/render/render_xlsx/fill_xlsx"
)

// Комплекс модулей для обработки файлов формата xlsx
type IRenderXlsx interface {
	fill_xlsx.IFillXlsx
}

type TRenderXlsx struct {
	fill_xlsx.IFillXlsx
}

func NewRenderXlsx() (IRenderXlsx, error) {
	t := &TRenderXlsx{}

	var err error
	t.IFillXlsx, err = fill_xlsx.NewFillXlsx()
	if err != nil {
		return nil, fmt.Errorf("NewRenderXlsx(): инициализация библиотеки заполнения, err=%w", err)
	}

	return t, nil
}
