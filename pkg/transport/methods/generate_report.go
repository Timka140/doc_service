package methods

import (
	"context"
	"encoding/json"
	"fmt"

	pb "projects/doc/doc_service/pkg/transport/protocol"
)

type TReport struct {
	Token  string  `json:"token"`   // Ключ авторизации
	Format string  `json:"formats"` // Формат отчета
	Code   string  `json:"code"`    // Код шаблона для документа
	Params TParams `json:"params"`  // Параметры объединения страницы

	Pack []byte `json:"pack"` // Данные для шаблона

	File TFile
}

// Задает параметры манипуляции над файлом
type TParams struct {
	NameFile   string `json:"name_file"`
	ConvertPDF bool   `json:"convert_pdf"` // Конвертация в PDF
	Rotation   bool   `json:"rotation"`    // Разворот страницы
	Join       bool   `json:"join"`        // Объединить в 1 файл
}

type TGenerateReports struct {
	Pack []byte // Данные или список данных для шаблона
}

// Входные параметры для создания документа запакованная в Pack
type TGenerateReportReqPack struct {
	Code   string                 `json:"code"`   // Тип шаблона
	Params map[string]interface{} `json:"params"` // Параметры шаблона
	Images []TImage               `json:"images"`
}

// Список изображений для замены
type TImage struct {
	Name  string `json:"name"`
	Image []byte `json:"image"`
}

// Результирующая структура ответа запакованная в Pack
type TGenerateReportRespPack struct {
	Files []*TFile `json:"files"`
}

type TFile struct {
	// Path     string //Путь где храниться файл
	Ext      string `json:"ext"`      // Формат отчета
	Name     string `json:"name"`     // Название
	FileData []byte `json:"fileData"` // Данные файлы
}

// GenerateReport() Создает отчет в зависимости от формата
func (t *TMethods) GenerateReport(val TGenerateReports) (res *TGenerateReportRespPack, err error) {
	if t.conn == nil {
		return nil, fmt.Errorf("TMethods.GenerateReport(): Соединение закрыто, err=%w", err)
	}

	pb, err := t.conn.GenerateReport(context.Background(), &pb.ReportReq{
		SrvAdr: &pb.ReportFormat{
			Token: t.token,
			Type:  "Render",
			Pack:  val.Pack,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("TMethods.GenerateReport(): Ошибка генерации отчета, err=%w", err)
	}

	res = &TGenerateReportRespPack{}
	err = json.Unmarshal(pb.SrvAdr.Pack, res)
	if err != nil {
		return nil, fmt.Errorf("TMethods.GenerateReport(): Ошибка структуры json, err=%w", err)
	}
	return res, nil
}
