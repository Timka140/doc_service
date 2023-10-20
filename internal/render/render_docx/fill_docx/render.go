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
		return fmt.Errorf("TFillDocx.RenderDocx(): неизвестный формат данных, err=%w", err)
	}
	//Устанавливаю параметры
	t.params = report.Params

	if pack.Code == "" {
		return fmt.Errorf("TFillDocx.RenderDocx(): код шаблона не задан")
	}

	if pack.Params == nil {
		return fmt.Errorf("TFillDocx.RenderDocx(): данные для шаблона не заданы")
	}

	file, err := t.select_template(pack.Code)
	if err != nil {
		return fmt.Errorf("TFillDocx.RenderDocx(): не удалось получить шаблон, err=%w", err)
	}

	fDocx, err := t.flow.Send(&interaction.TDocxIn{
		Template: file,
		Data:     pack.Params,
		Images:   pack.Images,
	})
	if err != nil {
		return fmt.Errorf("TFillDocx.RenderDocx(): отправка шаблона, err=%w", err)
	}

	if fDocx.Err != "" {
		return fmt.Errorf("TFillDocx.RenderDocx(): формирование, err=%v", fDocx.Err)
	}

	if fDocx.Data != nil {
		t.fileDocx.Write(fDocx.Data)
	}

	return nil
}
