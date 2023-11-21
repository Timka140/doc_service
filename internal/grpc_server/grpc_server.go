package grpc_server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"projects/doc/doc_service/internal/grpc_server/methods"
	pb "projects/doc/doc_service/pkg/transport/protocol"
)

type IGrpcServer interface {
}

type TGrpcServer struct {
	pb.ServiceServer
	methods.TMethods // Методы gRPC

	address  string
	listener net.Listener

	// cc       grpc.ClientConnInterface
	// mu sync.Mutex // protects routeNotes
}

// Создаем gRPC сервер
func NewServer(address string) (IGrpcServer, error) {
	var err error
	t := TGrpcServer{
		address: address,
	}

	//Инициализируем методы gRPC
	srv := methods.NewMethods()

	// Стартуем наш gRPC сервер для прослушивания tcp
	t.listener, err = net.Listen("tcp", t.address)
	if err != nil {
		return nil, fmt.Errorf("TGrpcServer.NewServer(): не удалось запустить сервер, err=%w", err)
	}

	maxSizeOption := grpc.MaxRecvMsgSize(32 * 10e6) //Устанавливаем максимальны размер сообщения 320MB
	sGRPC := grpc.NewServer(maxSizeOption)          // Подымаем gRPC сервер
	pb.RegisterServiceServer(sGRPC, srv)            // Регистрируем методы gRPC

	// Регистрация службы ответов на сервере gRPC.
	reflection.Register(sGRPC)
	if err := sGRPC.Serve(t.listener); err != nil {
		return nil, fmt.Errorf("TGrpcServer.NewServer(): не удалось зарегистрировать службы, err=%w", err)
	}

	return &t, nil
}
