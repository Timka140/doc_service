package docx

import (
	"bytes"

	"projects/doc/doc_service/internal/docx_service"
	"projects/doc/doc_service/internal/docx_service/interaction"
	"projects/doc/doc_service/pkg/transport/methods"
	"projects/doc/doc_service/pkg/types"
)

type tFillDocx struct {
	fileDocx *bytes.Buffer
	params   methods.TParams
	flow     interaction.IFlowDocx
}

func New() (types.ITemplateRender, error) {
	t := &tFillDocx{
		fileDocx: bytes.NewBuffer(nil),
		flow:     docx_service.FlowDocx,
	}

	return t, nil
}
