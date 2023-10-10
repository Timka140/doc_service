package methods

import (
	"context"

	pb "projects/doc/doc_service/pkg/transport/protocol"
)

// Методы поддерживаемые сервисом
type IMethods interface {
	GenerateReport(context.Context, *pb.ReportReq) (*pb.ReportResp, error) // Запускает процесс генерации файла на основе данных
	Ping(context.Context, *pb.PingReq) (*pb.PingResp, error)               // Проверка соединения
}
type TMethods struct {
	pb.ServiceServer
	// cc grpc.ClientConnInterface
}

func NewMethods() *TMethods {
	t := TMethods{}
	return &t
}
