package group_files

import (
	"fmt"
	"projects/doc/doc_service/pkg/transport/methods"
)

// Комплекс модулей отвечающий за создание различных типов документов docx формата
type IGroupFiles interface {
	// Инициализация генерации документов из нескольких шаблонов
	SendGroupFile() (IGroupFileSend, error)
}

type TGroupFiles struct {
	methods         methods.IMethods
	group_file_send IGroupFileSend
}

func NewGroupFiles(methods methods.IMethods) IGroupFiles {
	t := &TGroupFiles{
		methods:         methods,
		group_file_send: NewGroupFileSend(methods),
	}

	return t
}

func (t *TGroupFiles) SendGroupFile() (IGroupFileSend, error) {
	if t.group_file_send == nil {
		return nil, fmt.Errorf("TGroupFiles.SendGroupFile(): метод не инициализирован")
	}
	return t.group_file_send, nil
}
