package methods

import (
	"context"
	"encoding/json"
	"fmt"

	pb "projects/doc/doc_service/pkg/transport/protocol"
)

type TGenerateReportGroup struct {
	ReportFiles []*TReport `json:"report_files"` // Список шаблонов
	// MargeName   string     `json:"marge_name"`   // Название соединенного файла
}

// GenerateReport() Создает отчет в зависимости от формата
func (t *TMethods) GenerateReportGroup(val TGenerateReportGroup) (res *TGenerateReportRespPack, err error) {
	pack, err := json.Marshal(val)
	if err != nil {
		return nil, fmt.Errorf("TMethods.GenerateReportGroup() ошибка упаковки данных, err=%w", err)
	}

	pb, err := t.conn.GenerateReport(context.Background(), &pb.ReportReq{
		SrvAdr: &pb.ReportFormat{
			Type: "RenderGroup",
			Pack: pack,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("TMethods.GenerateReport(): Ошибка генерации отчета, err=%w", err)
	}

	res = &TGenerateReportRespPack{}
	err = json.Unmarshal(pb.SrvAdr.Pack, res)
	if err != nil {
		return nil, fmt.Errorf("TMethods.GenerateReport(): Ошибка структуры json, err=%w", err)
	}
	return res, nil
}
