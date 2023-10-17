package render

import (
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/internal/convert_pdf"
	"projects/doc/doc_service/internal/render/doc_one"
	"projects/doc/doc_service/internal/render/render_csv"
	"projects/doc/doc_service/internal/render/render_docx"
	"projects/doc/doc_service/internal/render/render_xlsx"

	"projects/doc/doc_service/pkg/transport/methods"
	pb "projects/doc/doc_service/pkg/transport/protocol"
)

/*
Модуль render отвечает за формирование отчета в различных форматах
*/

type IRender interface {
	SelectRender(SrvAdr *pb.ReportFormat) (result []byte, err error)
}

type TRender struct {
	docx   render_docx.IRenderDocx // Формирование документа docx
	xlsx   render_xlsx.IRenderXlsx // Формирование документа xlsx
	csv    render_csv.IRenderCsv   // Формирование документа csv
	docOne doc_one.IDocOne         //Формирование документов через doc_one

	convert_pdf convert_pdf.IConvertPDF //Конвертация и обработка PDF
}

// Инициализация движков формирования отчета
func NewRender() (IRender, error) {
	t := &TRender{}
	var err error

	// t.docOne, err = doc_one.NewDocOne()
	// if err != nil {
	// 	return nil, fmt.Errorf("render.NewRender(): ошибка инициализации doc_one: err =%w", err)
	// }

	t.docx, err = render_docx.NewRenderDocx()
	if err != nil {
		return nil, fmt.Errorf("render.NewRender(): ошибка инициализации render_docx: err =%w", err)
	}

	t.xlsx, err = render_xlsx.NewRenderXlsx()
	if err != nil {
		return nil, fmt.Errorf("render.NewRender(): ошибка инициализации render_xlsx: err =%w", err)
	}

	t.csv, err = render_csv.NewRenderCsv()
	if err != nil {
		return nil, fmt.Errorf("render.NewRender(): ошибка инициализации render_csv: err =%w", err)
	}

	t.convert_pdf, err = convert_pdf.NewConvertPDF()
	if err != nil {
		return nil, fmt.Errorf("render.NewRender(): ошибка инициализации convert_pdf: err =%w", err)
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
	// if SrvAdr.Type != "Render" {
	// 	return nil, nil
	// }

	reports, err := t.unpackReports(SrvAdr)
	if err != nil {
		return nil, fmt.Errorf("TRender.SelectRender(): инициация, err=%w", err)
	}

	for _, report := range reports.ReportFiles {
		switch report.Format {
		case "doc_one":
			// err = t.docOne.RenderDocOne(report)
			// if err != nil {
			// 	return nil, fmt.Errorf("TRender.SelectRender(): выполнение doc_one, err=%w", err)
			// }

			// report.File = t.docOne.GetDocument()
		case "docx":
			err = t.docx.RenderDocx(report)
			if err != nil {
				return nil, fmt.Errorf("TRender.SelectRender(): выполнение docx, err=%w", err)
			}

			report.File = t.docx.GetDocument()

		case "xlsx":
			err = t.xlsx.RenderXlsx(report)
			if err != nil {
				return nil, fmt.Errorf("TRender.SelectRender(): выполнение xlsx, err=%w", err)
			}

			report.File = t.xlsx.GetDocument()
		case "csv":
		default:
			return nil, fmt.Errorf("TRender.SelectRender(): неизвестный формат отчета")
		}
	}

	err = t.convert_pdf.SetData(reports)
	if err != nil {
		return nil, fmt.Errorf("TRender.SelectRender(): установка данных, err=%w", err)
	}

	err = t.convert_pdf.ConvertPDF()
	if err != nil {
		return nil, fmt.Errorf("TRender.SelectRender(): конвертация, err=%w", err)
	}

	pack, err := t.packReports(reports)
	if err != nil {
		return nil, fmt.Errorf("TRender.SelectRender(): отправка, err=%w", err)
	}

	return pack, nil
}
