package methods_pdf

import (
	"fmt"

	"projects/doc/doc_service/pkg/transport/methods"
)

type IMethodsPDF interface {
	SetDoc(doc *methods.TReport) error               //Устанавливает данные
	CrateCatalog(catalog string) error               // Создать каталог
	Convert() error                                  // Привести к формату pdf
	Rotate() error                                   //Повернуть pdf на 90 градусов
	MargePDF(in *methods.TGenerateReportGroup) error //Объединяет файлы

	CreateFile(name string) error //Создать файл
	RemoveCatalog() error         //Удалить каталог
	GetPathFile() string          // Возвращает путь файла

	AddMargeList()                                                  // Добавляет в список для объединения
	RemoveMargeFile(in []*methods.TReport) (out []*methods.TReport) //Убирает файлы которые были объеденины
}
type TMethodsPDF struct {
	file_path    string
	catalog_path string
	out_pdf      string

	marge_files []string
	doc         *methods.TReport
}

func NewMethodsPDF(out_pdf string) (IMethodsPDF, error) {
	t := &TMethodsPDF{
		out_pdf: out_pdf,
	}
	return t, nil
}

func (t *TMethodsPDF) SetDoc(doc *methods.TReport) error {
	if doc == nil {
		return fmt.Errorf("TMethodsPDF.SetDoc(): документ не задан")
	}

	t.doc = doc
	return nil
}
