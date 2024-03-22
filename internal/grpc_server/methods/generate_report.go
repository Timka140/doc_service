package methods

import (
	"context"
	"fmt"

	"projects/doc/doc_service/internal/grpc_server/grpc_sessions"
	"projects/doc/doc_service/internal/render"

	pb "projects/doc/doc_service/pkg/transport/protocol"
)

// GenerateReport Запускает процесс формирования файла.
func (t *TMethods) GenerateReport(ctx context.Context, in *pb.ReportReq) (out *pb.ReportResp, err error) {
	ses := grpc_sessions.Ses.GetSes(in.SrvAdr.Token)
	if ses == nil {
		return nil, fmt.Errorf("GenerateReport(): Сервис неавторизованн")
	}
	if !ses.Authorization() {
		return nil, fmt.Errorf("GenerateReport(): Сервис неавторизованн")
	}

	render, err := render.NewRender()
	if err != nil {
		return nil, fmt.Errorf("TMethods.GenerateReport(): не удалось сформировать render.NewRender(): err = %w", err)
	}

	// var pack methods.TGenerateReportRespPack
	result, err := render.SelectRender(in.SrvAdr)
	if err != nil {
		return nil, fmt.Errorf("TMethods.GenerateReport(): не удалось выбрать render.SelectRender(): err = %w", err)
	}

	// pack.FileData = result

	in.SrvAdr.Pack = result

	return &pb.ReportResp{SrvAdr: in.SrvAdr}, nil
}
