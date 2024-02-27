package connect

import (
	"context"
	"log"
	pb "projects/doc/doc_service/pkg/transport/protocol"
	"time"
)

func (t *TConnect) reconnect() {
	var err error
	defer func() {
		if err == nil {
			return
		}
		err = t.Close()
		if err != nil {
			log.Printf("connect.reconnect(): Закрытие соединения, err=%v", err)
		}
	}()
	err = t.create(t.info)
	if err != nil {
		log.Printf("connect.ping(): Восстановление соединения, err=%v", err)
		return
	}
}

func (t *TConnect) listenPing() {
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)

			resp, err := t.conn.Ping(context.Background(), &pb.PingReq{
				SrvPing: &pb.ServerPing{Sid: t.sid, Tm: time.Now().UnixMilli()},
			})
			if err != nil {
				log.Printf("connect.ping(): опрос сервера, err=%v", err)
				t.reconnect() //Запуск переподключения
			}
			start := time.UnixMilli(resp.SrvPing.Tm)
			end := time.Now()
			ms := end.Sub(start).Milliseconds()
			t.ping = ms
		}
	}()
}
