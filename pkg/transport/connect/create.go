package connect

import (
	"context"
	"encoding/json"
	"fmt"
	pb "projects/doc/doc_service/pkg/transport/protocol"
)

type TCreate struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (t *TConnect) create(in *TCreate) error {
	pack, _ := json.Marshal(in)
	_, err := t.conn.CreateSrv(context.Background(), &pb.CreateSrvReq{Srv: &pb.CreateService{Sid: t.sid, Pack: pack}})
	if err != nil {
		return fmt.Errorf("connect.info(): отправка информации о сервисе")
	}
	return nil
}
