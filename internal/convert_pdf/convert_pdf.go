package convert_pdf

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	"projects/doc/doc_service/internal/convert_pdf/methods_pdf"
	"projects/doc/doc_service/pkg/transport/methods"
)

type IConvertPDF interface {
	SetData(in *methods.TGenerateReportGroup) error
	ConvertPDF() error
}

type TConvertPDF struct {
	out_path string
	files    *methods.TGenerateReportGroup
	methods  methods_pdf.IMethodsPDF
}

func NewConvertPDF() (IConvertPDF, error) {
	out_pdf := os.Getenv("OUT_PDF")

	if out_pdf == "" {
		return nil, fmt.Errorf("NewConvertPDF(): переменная env не задана OUT_PDF")
	}

	var err error
	t := &TConvertPDF{
		out_path: out_pdf,
	}

	t.methods, err = methods_pdf.NewMethodsPDF(out_pdf)
	if err != nil {
		return nil, fmt.Errorf("NewConvertPDF(): инициализация методов, err=%w", err)
	}

	return t, nil
}

func (t *TConvertPDF) SetData(in *methods.TGenerateReportGroup) error {
	if in == nil {
		return fmt.Errorf("SetData(): установка данных")
	}
	t.files = in
	return nil
}

func (t *TConvertPDF) ConvertPDF() error {
	var err error
	catalog := uuid.NewString()
	err = t.methods.CrateCatalog(catalog)
	if err != nil {
		return fmt.Errorf("ConvertPDF(): создание каталога, err=%w", err)
	}

	for _, doc := range t.files.ReportFiles {
		name := fmt.Sprintf("%v.%v", uuid.NewString(), doc.File.Ext)

		err = t.methods.SetDoc(doc)
		if err != nil {
			return fmt.Errorf("ConvertPDF(): установка документа, err=%w", err)
		}

		err = t.methods.CreateFile(name)
		if err != nil {
			return fmt.Errorf("ConvertPDF(): создание документа, err=%w", err)
		}

		err = t.methods.Convert()
		if err != nil {
			return fmt.Errorf("ConvertPDF(): конвертация документа, err=%w", err)
		}

		err = t.methods.Rotate()
		if err != nil {
			return fmt.Errorf("ConvertPDF(): поворот документа, err=%w", err)
		}

		t.methods.AddMargeList()
	}

	err = t.methods.MargePDF(t.files)
	if err != nil {
		return fmt.Errorf("ConvertPDF(): объединение документа, err=%w", err)
	}

	t.files.ReportFiles = t.methods.RemoveMargeFile(t.files.ReportFiles)

	err = t.methods.RemoveCatalog()
	if err != nil {
		return fmt.Errorf("ConvertPDF(): удаление документа, err=%w", err)
	}
	return nil
}
