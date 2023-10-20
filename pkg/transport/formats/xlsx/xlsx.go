package xlsx

import "projects/doc/doc_service/pkg/transport/methods"

// Комплекс модулей отвечающий за создание различных типов документов docx формата
type IXlsx interface {
	// code - принимает ключ шаблона для генерации документа, params - принимает параметры для шаблона возвращает файл в бинарном виде
	XlsxPerform(code string, params methods.TParams, data map[string]interface{}) (res *methods.TFile, err error)
}

type TXlsx struct {
	methods methods.IMethods
}

func NewXlsx(methods methods.IMethods) IXlsx {
	t := &TXlsx{
		methods: methods,
	}

	return t
}
