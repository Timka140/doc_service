package transport

import (
	"fmt"

	"projects/doc/doc_service/pkg/transport/connect"
	"projects/doc/doc_service/pkg/transport/formats"
	"projects/doc/doc_service/pkg/transport/methods"
)

/*
Модуль отвечает за установку соединения и его закрытие
Передача данных осуществляется через gRPC
Описание протокола передачи данных находиться в папке protocol
*/

// Определение интерфейса ITransport
type ITransport interface {
	// connect.IConnect // Даем доступ к установке соединения
	// methods.IMethods // Даем доступ к методам протокола

	formats.IFormats
}

// Определение структуры TTransport
type TTransport struct {
	conn    connect.IConnect
	methods methods.IMethods
	formats.IFormats
}

// Функция NewTransport создает объект Transport и возвращает его
func NewTransport(address string) (ITransport, error) {
	var err error
	t := &TTransport{}

	t.conn = connect.NewConnect(address)

	err = t.conn.Open()
	if err != nil {
		return nil, fmt.Errorf("NewTransport(): открытие соединения err=%w", err)
	}

	t.methods, err = methods.NewMethods(t.conn.GetConn())
	if err != nil {
		return nil, fmt.Errorf("NewTransport(): ошибка соединения, err=%w", err)
	}

	t.IFormats = formats.NewFormats(t.methods)

	return t, nil
}
