package docx

import "projects/doc/doc_service/pkg/transport/methods"

// Комплекс модулей отвечающий за создание различных типов документов docx формата
type IDocx interface {
	// code - принимает ключ шаблона для генерации документа, params - принимает параметры для шаблона возвращает файл в бинарном виде
	DocxPerform(code string, params methods.TParams, data map[string]interface{}) (res *methods.TFile, err error)
}

type TDocx struct {
	methods methods.IMethods
}

func NewDocx(methods methods.IMethods) IDocx {
	t := &TDocx{
		methods: methods,
	}

	return t
}
