package methods

import (
	"context"
	"encoding/json"

	"projects/doc/doc_service/internal/services"
	"projects/doc/doc_service/pkg/transport/connect"
	pb "projects/doc/doc_service/pkg/transport/protocol"
)

// Info подключение нового микросервиса.
func (t *TMethods) CreateSrv(ctx context.Context, in *pb.CreateSrvReq) (out *pb.CreateSrvResp, err error) {
	data := &connect.TCreate{}
	_ = json.Unmarshal(in.Srv.Pack, data)

	srv := services.NewService(&services.TInServices{
		Sid:    in.Srv.Sid,
		Create: data,
	})

	services.Services.Add(in.Srv.Sid, srv)
	return &pb.CreateSrvResp{Srv: &pb.CreateService{}}, nil
}
