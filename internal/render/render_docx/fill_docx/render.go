package fill_docx

import (
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/internal/docx_service/interaction"
	"projects/doc/doc_service/pkg/transport/methods"
)

// Создает файл по шаблону docx
func (t *TFillDocx) RenderDocx(report *methods.TReport) (err error) {
	var pack methods.TGenerateReportReqPack
	err = json.Unmarshal(report.Pack, &pack)
	if err != nil {
		return fmt.Errorf("fill_docx.NewFillDocx(): неизвестный формат данных, err=%w", err)
	}
	//Устанавливаю параметры
	t.params = report.Params

	if pack.Code == "" {
		return fmt.Errorf("fill_docx.NewFillDocx(): код шаблона не задан")
	}

	if pack.Params == nil {
		return fmt.Errorf("fill_docx.NewFillDocx(): данные для шаблона не заданы")
	}

	file, err := t.select_template(pack.Code)
	if err != nil {
		return fmt.Errorf("fill_docx.NewFillDocx(): не удалось получить шаблон, err=%w", err)
	}

	fDocx, err := t.flow.Send(&interaction.TDocxIn{
		Template: file,
		Data:     pack.Params,
	})
	if err != nil {
		return fmt.Errorf("fill_docx.NewFillDocx(): отправка шаблона, err=%w", err)
	}

	if fDocx.Err != nil {
		return fmt.Errorf("fill_docx.NewFillDocx(): формирование, err=%v", fDocx.Err)
	}

	if fDocx.Data != nil {
		t.fileDocx.Write(fDocx.Data)
	}

	return nil
}
