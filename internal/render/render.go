package render

import (
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/internal/monitor"
	"projects/doc/doc_service/internal/render/csv"
	"projects/doc/doc_service/internal/render/docx"
	"projects/doc/doc_service/internal/render/pdf"
	"projects/doc/doc_service/internal/render/xlsx"

	"projects/doc/doc_service/pkg/cons"
	"projects/doc/doc_service/pkg/transport/methods"
	pb "projects/doc/doc_service/pkg/transport/protocol"
	"projects/doc/doc_service/pkg/types"
)

/*
Модуль render отвечает за формирование отчета в различных форматах
*/

type IRender interface {
	SelectRender(SrvAdr *pb.ReportFormat) (result []byte, err error)
}

type TRender struct {
	docx types.ITemplateRender // Формирование документа docx
	xlsx types.ITemplateRender // Формирование документа xlsx
	csv  csv.IRenderCsv        // Формирование документа csv
	pdf  types.ITemplateRender //Формирование и обработка PDF
	// docOne doc_one.IDocOne         //Формирование документов через doc_one

}

// Инициализация движков формирования отчета
func NewRender() (IRender, error) {
	t := &TRender{}
	var err error

	// t.docOne, err = doc_one.NewDocOne()
	// if err != nil {
	// 	return nil, fmt.Errorf("render.NewRender(): ошибка инициализации doc_one: err =%w", err)
	// }

	t.docx, err = docx.New()
	if err != nil {
		return nil, fmt.Errorf("render.NewRender(): ошибка инициализации docx: err =%w", err)
	}

	t.xlsx, err = xlsx.New()
	if err != nil {
		return nil, fmt.Errorf("render.NewRender(): ошибка инициализации render_xlsx: err =%w", err)
	}

	t.csv, err = csv.NewRenderCsv()
	if err != nil {
		return nil, fmt.Errorf("render.NewRender(): ошибка инициализации csv: err =%w", err)
	}

	t.pdf, err = pdf.New()
	if err != nil {
		return nil, fmt.Errorf("render.NewRender(): ошибка инициализации pdf: err =%w", err)
	}

	return t, nil
}

func (t *TRender) unpackReports(SrvAdr *pb.ReportFormat) (result *methods.TGenerateReportGroup, err error) {
	result = &methods.TGenerateReportGroup{}
	err = json.Unmarshal(SrvAdr.Pack, result)
	if err != nil {
		return nil, fmt.Errorf("TRender.SelectRender(): чтение входящего пакета, err=%w", err)
	}
	return result, nil
}

func (t *TRender) packReports(reports *methods.TGenerateReportGroup) (pack []byte, err error) {
	result := methods.TGenerateReportRespPack{}

	lFiles := len(reports.ReportFiles)
	result.Files = make([]*methods.TFile, lFiles)

	for ind, report := range reports.ReportFiles {
		result.Files[ind] = &report.File
	}

	pack, err = json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("TRender.packReports(): упаковка пакета, err=%w", err)
	}
	return pack, nil
}

func (t *TRender) SelectRender(SrvAdr *pb.ReportFormat) (result []byte, err error) {
	reports, err := t.unpackReports(SrvAdr)
	if err != nil {
		return nil, fmt.Errorf("TRender.SelectRender(): инициация, err=%w", err)
	}

	m := monitor.Monitor()

	for _, report := range reports.ReportFiles {
		if m != nil {
			m.Add(report.Format)
		}
		switch report.Format {
		case cons.CExtDocOne:
		case cons.CExtDocx:
			err = t.docx.Render(report)
			if err != nil {
				return nil, fmt.Errorf("TRender.SelectRender(): выполнение docx, err=%w", err)
			}
			report.File = t.docx.GetDocument()

		case cons.CExtXlsx:
			err = t.xlsx.Render(report)
			if err != nil {
				return nil, fmt.Errorf("TRender.SelectRender(): выполнение xlsx, err=%w", err)
			}

			report.File = t.xlsx.GetDocument()
		case cons.CExtCsv:
		case cons.CExtPdf:
			err = t.pdf.Render(report)
			if err != nil {
				return nil, fmt.Errorf("TRender.SelectRender(): выполнение pdf, err=%w", err)
			}

			report.File = t.pdf.GetDocument()
		default:
			return nil, fmt.Errorf("TRender.SelectRender(): неизвестный формат отчета")
		}
	}

	pack, err := t.packReports(reports)
	if err != nil {
		return nil, fmt.Errorf("TRender.SelectRender(): отправка, err=%w", err)
	}

	return pack, nil
}
