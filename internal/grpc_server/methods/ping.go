package methods

import (
	"context"
	"time"

	pb "projects/doc/doc_service/pkg/transport/protocol"
)

// Ping проверяет соединение и задержку.
func (t *TMethods) Ping(ctx context.Context, in *pb.PingReq) (out *pb.PingResp, err error) {
	in.SrvPing.Tm = time.Now().UnixMilli()
	return &pb.PingResp{SrvPing: in.SrvPing}, nil
}
