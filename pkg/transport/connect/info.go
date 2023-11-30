package connect

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	pb "projects/doc/doc_service/pkg/transport/protocol"
// )

type TInfo struct {
}

// func (t *TConnect) info(in *TInfo) error {
// 	pack, _ := json.Marshal(in)
// 	_, err := t.conn.Info(context.Background(), &pb.InfoReq{
// 		SrvInfo: &pb.Info{Sid: t.sid, Pack: pack},
// 	})
// 	if err != nil {
// 		return fmt.Errorf("connect.info(): отправка информации о сервисе")
// 	}
// 	return nil
// }
