package fill_pdf

import (
	"fmt"

	"projects/doc/doc_service/pkg/transport/methods"
)

// Создает файл по шаблону pdf
func (t *TFillPdf) RenderPdf(report *methods.TReport) (err error) {
	fPdf, err := t.flow.Send(report.Pack)
	if err != nil {
		return fmt.Errorf("TFillPdf.RenderPdf(): отправка шаблона, err=%w", err)
	}

	if fPdf.Err != "" {
		return fmt.Errorf("TFillPdf.RenderPdf(): формирование, err=%v", fPdf.Err)
	}

	if fPdf.Data != nil {
		t.filePdf.Write(fPdf.Data)
	}

	return nil
}
