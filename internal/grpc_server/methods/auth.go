package methods

import (
	"context"
	"fmt"

	"projects/doc/doc_service/internal/grpc_server/grpc_sessions"
	pb "projects/doc/doc_service/pkg/transport/protocol"

	"github.com/google/uuid"
)

// Auth проверка аккаунта.
func (t *TMethods) Auth(ctx context.Context, in *pb.AuthReq) (out *pb.AuthResp, err error) {
	ses, err := grpc_sessions.NewSession(in.SrvAuth.Key)
	if err != nil {
		return &pb.AuthResp{SrvAuth: in.SrvAuth}, fmt.Errorf("Auth(): авторизация, err=%w", err)
	}

	token := uuid.NewString()
	err = grpc_sessions.Ses.Add(token, ses)
	if err != nil {
		return nil, fmt.Errorf("Auth(): авторизация, err=%w", err)
	}
	in.SrvAuth.Token = token

	return &pb.AuthResp{SrvAuth: in.SrvAuth}, nil
}
