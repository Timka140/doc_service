package connect

import (
	"context"
	"log"
	pb "projects/doc/doc_service/pkg/transport/protocol"
	"time"
)

func (t *TConnect) listenPing() {
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			start := time.Now()
			resp, err := t.conn.Ping(context.Background(), &pb.PingReq{
				SrvPing: &pb.ServerPing{Sid: t.sid, Tm: time.Now().UnixMilli()},
			})
			if err != nil {
				log.Printf("connect.ping(): опрос сервера, err=%v", err)
			}
			end := time.UnixMilli(resp.SrvPing.Tm)
			ms := end.Sub(start).Milliseconds()
			t.ping = ms
		}
	}()
}
