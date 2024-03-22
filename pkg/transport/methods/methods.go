package methods

import (
	"fmt"

	pb "projects/doc/doc_service/pkg/transport/protocol"
)

type IMethods interface {
	// Token() string
	GenerateReport(val TGenerateReports) (res *TGenerateReportRespPack, err error) //Запускает генерацию отчета
}

type TMethods struct {
	token string
	conn  pb.ServiceClient // Клиент для взаимодействия с сервисом
}

// Установка соединения
func NewMethods(conn pb.ServiceClient, token string) (IMethods, error) {
	if conn == nil {
		return nil, fmt.Errorf("NewMethods(): соединение не установлено")
	}
	t := &TMethods{
		token: token,
		conn:  conn,
	}
	return t, nil
}

// func (t *TMethods) Token() string {
// 	return t.token
// }
