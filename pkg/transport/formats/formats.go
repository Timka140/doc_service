package formats

import (
	formats_doc_one "projects/doc/doc_service/pkg/transport/formats/doc_one"
	formats_docx "projects/doc/doc_service/pkg/transport/formats/docx"
	"projects/doc/doc_service/pkg/transport/formats/group_files"
	formats_xlsx "projects/doc/doc_service/pkg/transport/formats/xlsx"
	"projects/doc/doc_service/pkg/transport/methods"
)

// Набор методов возвращающие различные типы файлов
type IFormats interface {
	formats_docx.IDocx
	formats_xlsx.IXlsx
	formats_doc_one.IDocOne
	group_files.IGroupFiles
}

type TFormats struct {
	methods methods.IMethods
	formats_docx.IDocx
	formats_xlsx.IXlsx
	formats_doc_one.IDocOne
	group_files.IGroupFiles
}

func NewFormats(methods methods.IMethods) IFormats {
	formats := &TFormats{
		methods: methods,
	}

	formats.IDocx = formats_docx.NewDocx(formats.methods)            //Методы для отправки файлов docx
	formats.IXlsx = formats_xlsx.NewXlsx(formats.methods)            //Методы для отправки файлов docx
	formats.IDocOne = formats_doc_one.NewDocOne(formats.methods)     //Методы для отправки файлов doc_one
	formats.IGroupFiles = group_files.NewGroupFiles(formats.methods) //Методы для отправки пакета файлов

	return formats
}
