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
	Token() string             //Ключ сессии
}
type TConnect struct {
	key     string
	token   string
	address string // Адрес, с которым будет установлено соединение
	ping    int64

	info  *TCreate
	cConn *grpc.ClientConn // Клиентское подключение grpc
	conn  pb.ServiceClient // Клиент для взаимодействия с сервисом
}

// Инициализация нового соединения
func NewConnect(address, key string, info *TCreate) (IConnect, error) {
	if key == "" {
		return nil, fmt.Errorf("NewConnect(): Ключ авторизации не указан")
	}

	if info == nil {
		return nil, fmt.Errorf("NewConnect(): информация о сервисе не задана")
	}

	if info.Name == "" {
		return nil, fmt.Errorf("NewConnect(): название сервиса не указано")
	}

	t := &TConnect{
		// token:   uuid.NewString(),
		key:     key,
		address: address,
		info:    info,
	}
	return t, nil
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

	err = t.auth()
	if err != nil {
		return fmt.Errorf("TServer.Open(): Авторизация: %w", err) // Если произошла ошибка, возвращаем сообщение с ошибкой
	}
	t.create(t.info)

	t.listenPing() //инициализируем пинг

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

func (t *TConnect) Token() string {
	return t.token // Ключ авторизации
}
