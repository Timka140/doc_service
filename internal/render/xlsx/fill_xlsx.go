package xlsx

import (
	"bytes"

	"projects/doc/doc_service/internal/xlsx_service"
	"projects/doc/doc_service/internal/xlsx_service/interaction"
	"projects/doc/doc_service/pkg/transport/methods"
	"projects/doc/doc_service/pkg/types"

	"github.com/xuri/excelize/v2"
)

type tFillXlsx struct {
	fXlsx *excelize.File

	fileXlsx *bytes.Buffer
	params   methods.TParams

	flow interaction.IFlowXlsx
}

func New() (types.ITemplateRender, error) {

	t := &tFillXlsx{

		fileXlsx: bytes.NewBuffer(nil),
		flow:     xlsx_service.FlowXlsx,
	}

	return t, nil
}
