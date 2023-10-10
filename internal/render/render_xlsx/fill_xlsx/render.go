package fill_xlsx

import (
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/pkg/transport/methods"
)

// Создает файл по шаблону docx
func (t *TFillXlsx) RenderXlsx(report *methods.TReport) (err error) {
	var pack methods.TGenerateReportReqPack
	err = json.Unmarshal(report.Pack, &pack)
	if err != nil {
		return fmt.Errorf("fill_docx.RenderXlsx(): неизвестный формат данных, err=%w", err)
	}
	//Устанавливаю параметры
	t.params = report.Params

	if pack.Code == "" {
		return fmt.Errorf("fill_docx.RenderXlsx(): код шаблона не задан")
	}

	if pack.Params == nil {
		return fmt.Errorf("fill_docx.RenderXlsx(): данные для шаблона не заданы")
	}

	var funcEdit string
	t.fXlsx, funcEdit, err = t.select_template(pack.Code)
	if err != nil {
		return fmt.Errorf("fill_docx.RenderXlsx(): не удалось получить шаблон, err=%w", err)
	}

	err = t.fill_template(funcEdit, pack.Params)
	if err != nil {
		return fmt.Errorf("fill_docx.RenderXlsx(): заполнение шаблона данными, err=%w", err)
	}

	err = t.fXlsx.Close()
	if err != nil {
		return fmt.Errorf("fill_docx.RenderXlsx(): закрытие файла docx, err=%w", err)
	}

	return nil
}
