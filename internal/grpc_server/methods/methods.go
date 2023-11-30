package methods

import (
	"context"

	pb "projects/doc/doc_service/pkg/transport/protocol"
)

// Методы поддерживаемые сервисом
type IMethods interface {
	GenerateReport(context.Context, *pb.ReportReq) (*pb.ReportResp, error)  // Запускает процесс генерации файла на основе данных
	Ping(context.Context, *pb.PingReq) (*pb.PingResp, error)                // Проверка соединения
	Info(context.Context, *pb.InfoReq) (*pb.InfoResp, error)                // Информация о микро сервисе
	CreateSrv(context.Context, *pb.CreateSrvReq) (*pb.CreateSrvResp, error) // Установка соединения
}
type TMethods struct {
	pb.ServiceServer
	// cc grpc.ClientConnInterface
}

func NewMethods() *TMethods {
	t := TMethods{}
	return &t
}
