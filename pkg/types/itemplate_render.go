package types

import "projects/doc/doc_service/pkg/transport/methods"

type ITemplateRender interface {
	// Render -- Создает файл по шаблону
	Render(report *methods.TReport) (err error)
	// ReadBytes -- Возвращает файл в байтовом виде
	ReadBytes() (data []byte, err error)
	// WriteToPath -- Записывает файл по нужному пути
	WriteToPath(path string) (err error)
	// GetDocument -- Возвращает структуру запакованного документа документа
	GetDocument() (file methods.TFile)
}
