package docx

import (
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/pkg/cons"
	"projects/doc/doc_service/pkg/transport/methods"
)

/*
code - принимает ключ шаблона для генерации документа
params - принимает параметры для шаблона
*/
func (t *TDocx) DocxPerform(code string, params methods.TParams, data map[string]interface{}) (res *methods.TFile, err error) {

	pack, err := json.Marshal(methods.TGenerateReportReqPack{
		Code:   code,
		Params: data,
	})

	if err != nil {
		return nil, fmt.Errorf("TDocx.DocxPerform() ошибка упаковки данных, err=%w", err)
	}

	pack_reports, err := json.Marshal(
		methods.TGenerateReportGroup{
			ReportFiles: []*methods.TReport{
				{
					Format: cons.CExtDocx,
					Code:   code,
					Pack:   pack,
					Params: params,
				},
			},
		})

	if err != nil {
		return nil, fmt.Errorf("TDocx.DocxPerform() ошибка упаковки массива данных, err=%w", err)
	}

	pack_res, err := t.methods.GenerateReport(methods.TGenerateReports{
		Pack: pack_reports,
	})

	if err != nil {
		return nil, fmt.Errorf("TDocx.DocxPerform() ошибка упаковки данных, err=%w", err)
	}

	if len(pack_res.Files) != 1 {
		return nil, fmt.Errorf("TDocx.DocxPerform() !!! Неожиданные данные, err=%w", err)
	}

	return pack_res.Files[0], nil
}
