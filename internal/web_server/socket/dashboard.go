package socket

import (
	"fmt"
	"log"
	"projects/doc/doc_service/internal/docx_service"
	"projects/doc/doc_service/internal/pdf_service"
	"projects/doc/doc_service/internal/xlsx_service"
	"projects/doc/doc_service/pkg/types"

	"github.com/google/uuid"
)

type TDashboard struct {
	data map[string]interface{}
	pid  string

	docx      docx_service.IDocxService
	docxWorks types.IWorkers

	pdf      pdf_service.IPdfService
	pdfWorks types.IWorkers

	xlsx      xlsx_service.IXlsxService
	xlsxWorks types.IWorkers
}

func newDashboardSocket(in *TSocketValue) (ISocket, error) {
	t := &TDashboard{
		data: in.Data,
		pid:  uuid.NewString(),
		docx: docx_service.DocxServices,
		pdf:  pdf_service.PdfServices,
		xlsx: xlsx_service.XlsxServices,
	}
	var err error
	t.docxWorks, err = t.docx.Workers()
	if err != nil {
		return nil, fmt.Errorf("TDashboard.Start(): docx works, err=%w", err)
	}
	t.pdfWorks, err = t.pdf.Workers()
	if err != nil {
		return nil, fmt.Errorf("TDashboard.Start(): pdf works, err=%w", err)
	}
	t.xlsxWorks, err = t.xlsx.Workers()
	if err != nil {
		return nil, fmt.Errorf("TDashboard.Start(): xlsx works, err=%w", err)
	}

	return t, nil
}

func (t *TDashboard) Start() error {

	execution, ok := t.data["execution"].(string)
	if !ok {
		return fmt.Errorf("TDashboard.Start(): не прочитано исполнение")
	}

	switch execution {
	case "init":
		t.data["docx_len"] = t.docxWorks.Len()
		t.data["pdf_len"] = t.pdfWorks.Len()
		t.data["xlsx_len"] = t.xlsxWorks.Len()
	case "info":

	}

	return nil
}

func (t *TDashboard) GetPid() string {
	return t.pid
}

func (t *TDashboard) Response() (map[string]interface{}, error) {
	return t.data, nil
}

func (t *TDashboard) Stop() error {

	return nil
}

func init() {
	err := constructors.Add("Dashboard", newDashboardSocket)
	if err != nil {
		log.Printf("Dashboard(): не удалось добавить в конструктор")
	}
}
