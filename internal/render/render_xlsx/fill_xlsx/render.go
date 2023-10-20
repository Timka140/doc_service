package fill_xlsx

import (
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/internal/xlsx_service/interaction"
	"projects/doc/doc_service/pkg/transport/methods"
)

// Создает файл по шаблону docx
func (t *TFillXlsx) RenderXlsx(report *methods.TReport) (err error) {
	var pack methods.TGenerateReportReqPack
	err = json.Unmarshal(report.Pack, &pack)
	if err != nil {
		return fmt.Errorf("TFillXlsx.RenderXlsx(): неизвестный формат данных, err=%w", err)
	}
	//Устанавливаю параметры
	t.params = report.Params

	if pack.Code == "" {
		return fmt.Errorf("TFillXlsx.RenderXlsx(): код шаблона не задан")
	}

	if pack.Params == nil {
		return fmt.Errorf("TFillXlsx.RenderXlsx(): данные для шаблона не заданы")
	}

	file, err := t.select_template(pack.Code)
	if err != nil {
		return fmt.Errorf("TFillXlsx.RenderXlsx(): не удалось получить шаблон, err=%w", err)
	}

	fXlsx, err := t.flow.Send(&interaction.TXlsxIn{
		Template: file,
		Data:     pack.Params,
		Images:   pack.Images,
	})
	if err != nil {
		return fmt.Errorf("TFillXlsx.RenderXlsx(): отправка шаблона, err=%w", err)
	}

	if fXlsx.Err != "" {
		return fmt.Errorf("TFillXlsx.RenderXlsx(): формирование, err=%v", fXlsx.Err)
	}

	if fXlsx.Data != nil {
		t.fileXlsx.Write(fXlsx.Data)
	}

	return nil
}
