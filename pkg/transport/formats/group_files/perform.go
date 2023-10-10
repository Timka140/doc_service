package group_files

import (
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/pkg/cons"
	formats_xlsx "projects/doc/doc_service/pkg/transport/formats/xlsx"
	"projects/doc/doc_service/pkg/transport/methods"
)

type IGroupFileSend interface {
	DocOnePerform(code string, params methods.TParams, data map[string][]string) error
	DocxPerform(code string, params methods.TParams, data map[string]interface{}) error
	XlsxPerform(code string, params methods.TParams, data formats_xlsx.TValuesRender) error

	Send() (res []*methods.TFile, err error)
}

type TGroupFileSend struct {
	methods methods.IMethods
	data    methods.TGenerateReportGroup
}

func NewGroupFileSend(methods methods.IMethods) IGroupFileSend {
	t := &TGroupFileSend{
		methods: methods,
	}
	return t
}

func (t *TGroupFileSend) DocOnePerform(code string, params methods.TParams, data map[string][]string) error {
	buf := make(map[string]interface{})

	for key, v := range data {
		buf[key] = v
	}

	err := t.addFile(cons.CExtDocOne, code, params, buf)
	if err != nil {
		return fmt.Errorf("TGroupFiles.DocOnePerform() ошибка формирования данных, err=%w", err)
	}
	return nil
}

func (t *TGroupFileSend) DocxPerform(code string, params methods.TParams, data map[string]interface{}) error {
	err := t.addFile(cons.CExtDocx, code, params, data)
	if err != nil {
		return fmt.Errorf("TGroupFiles.DocOnePerform() ошибка формирования данных, err=%w", err)
	}
	return nil
}

func (t *TGroupFileSend) XlsxPerform(code string, params methods.TParams, data formats_xlsx.TValuesRender) error {
	buf := make(map[string]interface{})

	buf["Data"] = data.Data
	buf["DataTable"] = data.DataTable
	buf["Pages"] = data.Pages

	err := t.addFile(cons.CExtXlsx, code, params, buf)
	if err != nil {
		return fmt.Errorf("TGroupFiles.DocOnePerform() ошибка формирования данных, err=%w", err)
	}
	return nil
}

func (t *TGroupFileSend) addFile(format, code string, params methods.TParams, data map[string]interface{}) error {

	if format == "" {
		return fmt.Errorf("TGroupFiles.AddFile() format не задан")
	}

	if code == "" {
		return fmt.Errorf("TGroupFiles.AddFile() code не задан")
	}

	if data == nil {
		return fmt.Errorf("TGroupFiles.AddFile() data не задан")
	}

	pack, err := json.Marshal(methods.TGenerateReportReqPack{
		Code:   code,
		Params: data,
	})

	if err != nil {
		return fmt.Errorf("TGroupFiles.AddFile() ошибка упаковки данных, err=%w", err)
	}

	t.data.ReportFiles = append(t.data.ReportFiles, &methods.TReport{
		Format: format,
		Code:   code,
		Pack:   pack,
		Params: params,
	})

	return nil
}

func (t *TGroupFileSend) Send() (res []*methods.TFile, err error) {

	pack_reports, err := json.Marshal(t.data)

	if err != nil {
		return nil, fmt.Errorf("TDocx.GroupFilePerform() ошибка упаковки массива данных, err=%w", err)
	}

	pack_res, err := t.methods.GenerateReport(methods.TGenerateReports{
		Pack: pack_reports,
	})

	if err != nil {
		return nil, fmt.Errorf("TDocx.GroupFilePerform() ошибка упаковки данных, err=%w", err)
	}

	return pack_res.Files, nil
}
