package formats

import (
	"projects/doc/doc_service/pkg/transport/formats/docx"
	"projects/doc/doc_service/pkg/transport/formats/pdf"
	"projects/doc/doc_service/pkg/transport/methods"
)

// Набор методов возвращающие различные типы файлов
type IFormats interface {
	docx.IDocx
	// xlsx.IXlsx
	pdf.IPdf
	// formats_doc_one.IDocOne
}

type TFormats struct {
	methods methods.IMethods
	docx.IDocx
	// xlsx.IXlsx
	pdf.IPdf
	// formats_doc_one.IDocOne
}

func NewFormats(methods methods.IMethods) IFormats {
	formats := &TFormats{
		methods: methods,
	}

	formats.IDocx = docx.NewDocx(formats.methods) //Методы для отправки файлов docx
	// formats.IXlsx = xlsx.NewXlsx(formats.methods)                //Методы для отправки файлов xlsx
	formats.IPdf = pdf.NewPdf(formats.methods) //Методы для отправки файлов pdf
	// formats.IDocOne = formats_doc_one.NewDocOne(formats.methods) //Методы для отправки файлов doc_one

	return formats
}
