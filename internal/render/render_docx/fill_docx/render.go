package fill_docx

import (
	"encoding/json"
	"fmt"
	"log"

	"projects/doc/doc_service/internal/docx_service"
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

	fDocx, err := docx_service.DocxServices.FillDocx()
	if err != nil {
		return fmt.Errorf("fill_docx.NewFillDocx(): заполнение шаблона данными, err=%w", err)
	}
	tmp, err := fDocx.Pack(&interaction.TDocxIn{
		Template: file,
		Data:     pack.Params,
	})
	if err != nil {
		return fmt.Errorf("fill_docx.NewFillDocx(): упаковка, err=%w", err)
	}

	err = fDocx.Send(tmp)
	if err != nil {
		return fmt.Errorf("fill_docx.NewFillDocx(): отправка шаблона, err=%w", err)
	}

	err = fDocx.Result(func(data *interaction.TDocxOut) {
		if data.Err != nil {
			log.Printf("fill_docx.NewFillDocx(): формирование документа, err=%v", data.Err)
			return
		}
		t.fileDocx.Write(data.Data)
	})
	if err != nil {
		return fmt.Errorf("fill_docx.NewFillDocx(): чтение результата, err=%w", err)
	}

	return nil
}
