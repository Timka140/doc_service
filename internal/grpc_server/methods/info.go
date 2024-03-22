package methods

import (
	"context"
	"fmt"

	"projects/doc/doc_service/internal/grpc_server/grpc_sessions"
	pb "projects/doc/doc_service/pkg/transport/protocol"
)

// Info проверяет соединение и задержку.
func (t *TMethods) Info(ctx context.Context, in *pb.InfoReq) (out *pb.InfoResp, err error) {
	var pack []byte

	ses := grpc_sessions.Ses.GetSes(in.SrvInfo.Token)
	if ses == nil {
		return nil, fmt.Errorf("Info(): Сервис неавторизованн")
	}
	if !ses.Authorization() {
		return nil, fmt.Errorf("Info(): Сервис неавторизованн")
	}

	// data := make(map[string]interface{})

	// func() {
	// 	defer func() {
	// 		pack, _ = json.Marshal(data)
	// 	}()

	// 	srv, err := services.Services.Get(in.SrvInfo.Sid)
	// 	if err != nil {
	// 		data["err"] = fmt.Errorf("Info(): обращение к микросервису, err=%w", err)
	// 		return
	// 	}

	// 	info := &connect.TInfo{}
	// 	_ = json.Unmarshal(in.SrvInfo.Pack, info)

	// 	srv.SetInfo(info)

	// }()

	return &pb.InfoResp{SrvInfo: &pb.Info{
		Pack: pack,
	}}, nil
}
