package xlsx

import (
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/pkg/cons"
	"projects/doc/doc_service/pkg/transport/methods"
)

// Структура документа
type TLine struct {
	StartPosition int
	Data          map[string]interface{}
}
type TPage struct {
	Name  string
	Lines []TLine
}

type TValuesRender struct {
	DataTable []map[string]interface{}
	Data      map[string]interface{}
	Pages     []TPage
}

func (t *TXlsx) XlsxPerform(code string, params methods.TParams, data map[string]interface{}) (*methods.TFile, error) {

	pack, err := json.Marshal(methods.TGenerateReportReqPack{
		Code:   code,
		Params: data,
	})

	if err != nil {
		return nil, fmt.Errorf("TDocx.XlsxPerform() ошибка упаковки данных, err=%w", err)
	}

	pack_reports, err := json.Marshal(
		methods.TGenerateReportGroup{
			ReportFiles: []*methods.TReport{
				{
					Format: cons.CExtXlsx,
					Code:   code,
					Pack:   pack,
					Params: params,
				},
			},
		})

	if err != nil {
		return nil, fmt.Errorf("TDocx.XlsxPerform() ошибка упаковки массива данных, err=%w", err)
	}

	pack_res, err := t.methods.GenerateReport(methods.TGenerateReports{
		Pack: pack_reports,
	})

	if err != nil {
		return nil, fmt.Errorf("TDocx.XlsxPerform() ошибка упаковки данных, err=%w", err)
	}

	if len(pack_res.Files) != 1 {
		return nil, fmt.Errorf("TDocx.XlsxPerform() !!! Неожиданные данные, err=%w", err)
	}

	return pack_res.Files[0], nil
}
