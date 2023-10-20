package pdf

import (
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/pkg/cons"
	"projects/doc/doc_service/pkg/transport/methods"
)

type TFiles struct {
	Files []*TFile `json:"files"`
}
type TFile struct {
	Params   TParams `json:"params"`
	Ext      string  `json:"ext"`
	Name     string  `json:"name"`
	FileData []byte  `json:"data"`
}
type TParams struct {
	// Join     bool `json:"join"`
	Rotation bool `json:"rotation"`
}

func (t *TPdf) PdfPerform(files TFiles) (*methods.TFile, error) {

	pack, err := json.Marshal(files)

	if err != nil {
		return nil, fmt.Errorf("TPdf.PdfPerform() ошибка упаковки данных, err=%w", err)
	}

	pack_reports, err := json.Marshal(
		methods.TGenerateReportGroup{
			ReportFiles: []*methods.TReport{
				{
					Format: cons.CExtPdf,
					Pack:   pack,
				},
			},
		})

	if err != nil {
		return nil, fmt.Errorf("TPdf.PdfPerform() ошибка упаковки массива данных, err=%w", err)
	}

	pack_res, err := t.methods.GenerateReport(methods.TGenerateReports{
		Pack: pack_reports,
	})

	if err != nil {
		return nil, fmt.Errorf("TPdf.PdfPerform() ошибка упаковки данных, err=%w", err)
	}

	if len(pack_res.Files) != 1 {
		return nil, fmt.Errorf("TPdf.PdfPerform() !!! Неожиданные данные, err=%w", err)
	}

	return pack_res.Files[0], nil
}
