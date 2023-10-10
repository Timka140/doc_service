package connect

import (
	"fmt"

	"google.golang.org/grpc/credentials/insecure"

	pb "projects/doc/doc_service/pkg/transport/protocol"

	grpc "google.golang.org/grpc"
)

type IConnect interface {
	Open() error               //Установка соединения
	Close() error              // Закрытие соединения
	GetConn() pb.ServiceClient // Получаем соединение
}
type TConnect struct {
	address string           // Адрес, с которым будет установлено соединение
	cConn   *grpc.ClientConn // Клиентское подключение grpc
	conn    pb.ServiceClient // Клиент для взаимодействия с сервисом
}

// Инициализация нового соединения
func NewConnect(address string) IConnect {
	t := &TConnect{
		address: address,
	}
	return t
}

// Open устанавливает соединение с адресом, указанным в поле address
func (t *TConnect) Open() error {
	var err error

	opts := grpc.WithTransportCredentials(insecure.NewCredentials()) // Используем не надежные учетные данные для установления соединения

	t.cConn, err = grpc.Dial(t.address, opts) // Устанавливаем клиентское подключение по указанному адресу
	if err != nil {
		return fmt.Errorf("TServer.Open(): Не удалось установить соединение: %w", err) // Если произошла ошибка, возвращаем сообщение с ошибкой
	}

	t.conn = pb.NewServiceClient(t.cConn) // Создаем клиента для взаимодействия со службой

	return nil // Успешное открытие соединения
}

// Close закрывает клиентское подключение
func (t *TConnect) Close() error {
	err := t.cConn.Close() // Закрываем клиентское подключение
	if err != nil {
		return fmt.Errorf("TServer.Close(): ошибка закрытия соединения, err=%w", err) // Если произошла ошибка, возвращаем сообщение с ошибкой
	}
	return nil // Успешное закрытие соединения
}

// GetConn дает доступ к соединению
func (t *TConnect) GetConn() pb.ServiceClient {
	return t.conn // Возвращаем соединение
}
