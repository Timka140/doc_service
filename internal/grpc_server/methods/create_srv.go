package methods

import (
	"context"
	"encoding/json"
	"fmt"

	"projects/doc/doc_service/internal/grpc_server/grpc_sessions"
	"projects/doc/doc_service/internal/services"
	"projects/doc/doc_service/pkg/transport/connect"
	pb "projects/doc/doc_service/pkg/transport/protocol"
)

// Info подключение нового микросервиса.
func (t *TMethods) CreateSrv(ctx context.Context, in *pb.CreateSrvReq) (out *pb.CreateSrvResp, err error) {
	ses := grpc_sessions.Ses.GetSes(in.Srv.Token)
	if ses == nil {
		return nil, fmt.Errorf("CreateSrv(): Сервис неавторизованн")
	}
	if !ses.Authorization() {
		return nil, fmt.Errorf("CreateSrv(): Сервис неавторизованн")
	}

	data := &connect.TCreate{}
	_ = json.Unmarshal(in.Srv.Pack, data)

	srv := services.NewService(&services.TInServices{
		Sid:    in.Srv.Token,
		Create: data,
	})

	services.Services.Add(in.Srv.Token, srv)
	return &pb.CreateSrvResp{Srv: &pb.CreateService{}}, nil
}
