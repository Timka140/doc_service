package connect

import (
	"context"
	"fmt"
	pb "projects/doc/doc_service/pkg/transport/protocol"
)

func (t *TConnect) auth() error {
	resp, err := t.conn.Auth(context.Background(), &pb.AuthReq{
		SrvAuth: &pb.ServerAuth{Key: t.key},
	})
	if err != nil {
		return fmt.Errorf("connect.auth(): авторизация, err=%w", err)
	}
	t.token = resp.SrvAuth.Token
	return nil
}
