package services

import (
	"projects/doc/doc_service/pkg/transport/connect"
	"time"
)

type TService struct {
	name    string
	comment string

	sid    string
	update time.Time

	ping int64 // в миллисекундах

	data map[string]interface{}
}

func NewService(in *TInServices) IService {
	t := &TService{
		data: make(map[string]interface{}),

		name:    in.Create.Name,
		comment: in.Create.Comment,
		sid:     in.Sid,
	}

	go func() {
		for {
			time.Sleep(1 * time.Minute)
			tm := time.Now()
			mt := tm.Sub(t.update).Minutes()
			if mt > 5 {
				Services.Delete(t.sid)
				return
			}
		}
	}()

	return t
}

func (t *TService) Name() string {
	return t.name
}
func (t *TService) Comment() string {
	return t.comment
}

func (t *TService) Ping() int64 {
	return t.ping
}

func (t *TService) SetPing(ping int64) {
	t.update = time.Now()
	t.ping = ping
}

func (t *TService) Info() map[string]interface{} {
	return t.data
}
func (t *TService) SetInfo(in *connect.TInfo) {

}

func (t *TService) Data() map[string]interface{} {
	return t.data
}
