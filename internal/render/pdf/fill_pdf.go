package pdf

import (
	"bytes"

	"projects/doc/doc_service/internal/pdf_service"
	"projects/doc/doc_service/internal/pdf_service/interaction"

	"projects/doc/doc_service/pkg/transport/methods"
	"projects/doc/doc_service/pkg/types"
)

type tFillPdf struct {
	filePdf *bytes.Buffer
	params  methods.TParams
	flow    interaction.IFlowPdf
}

func New() (types.ITemplateRender, error) {
	t := &tFillPdf{
		filePdf: bytes.NewBuffer(nil),
		flow:    pdf_service.FlowPdf,
	}

	return t, nil
}
