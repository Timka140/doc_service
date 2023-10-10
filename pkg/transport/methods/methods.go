package methods

import (
	"fmt"

	pb "projects/doc/doc_service/pkg/transport/protocol"
)

type IMethods interface {
	GenerateReport(val TGenerateReports) (res *TGenerateReportRespPack, err error) //Запускает генерацию отчета
}

type TMethods struct {
	conn pb.ServiceClient // Клиент для взаимодействия с сервисом
}

// Установка соединения
func NewMethods(conn pb.ServiceClient) (IMethods, error) {
	if conn == nil {
		return nil, fmt.Errorf("NewMethods(): соединение не установлено")
	}
	t := &TMethods{
		conn: conn,
	}
	return t, nil
}
