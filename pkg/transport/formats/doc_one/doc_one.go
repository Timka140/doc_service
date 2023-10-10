package formats_doc_one

import "projects/doc/doc_service/pkg/transport/methods"

// Комплекс модулей отвечающий за создание различных типов документов docx формата
type IDocOne interface {
	// code - принимает ключ шаблона для генерации документа, params - принимает параметры для шаблона возвращает файл в бинарном виде
	DocOnePerform(code string, params methods.TParams, data map[string]interface{}) (res *methods.TFile, err error)
}

type TDocOne struct {
	methods methods.IMethods
}

func NewDocOne(methods methods.IMethods) IDocOne {
	t := &TDocOne{
		methods: methods,
	}

	return t
}
