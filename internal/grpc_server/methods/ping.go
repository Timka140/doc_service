package methods

import (
	"context"
	"fmt"
	"time"

	"projects/doc/doc_service/internal/services"
	pb "projects/doc/doc_service/pkg/transport/protocol"
)

// Ping проверяет соединение и задержку.
func (t *TMethods) Ping(ctx context.Context, in *pb.PingReq) (out *pb.PingResp, err error) {
	// time.Sleep(time.Duration(rand.Intn(30)) * time.Millisecond)
	start := time.Now()
	end := time.UnixMilli(in.SrvPing.Tm)
	ms := start.Sub(end).Milliseconds()

	srv, err := services.Services.Get(in.SrvPing.Sid)
	if err != nil {
		return &pb.PingResp{SrvPing: in.SrvPing}, fmt.Errorf("Info(): обращение к микросервису, err=%w", err)
	}
	srv.SetPing(ms)

	in.SrvPing.Tm = end.UnixMilli()
	return &pb.PingResp{SrvPing: in.SrvPing}, nil
}
