package pdf

import "projects/doc/doc_service/pkg/transport/methods"

type IPdf interface {
	// files - принимает список файлов для преобразования в PDF
	PdfPerform(files TFiles) (*methods.TFile, error)
}

type TPdf struct {
	methods methods.IMethods
}

func NewPdf(methods methods.IMethods) IPdf {
	t := &TPdf{
		methods: methods,
	}

	return t
}
