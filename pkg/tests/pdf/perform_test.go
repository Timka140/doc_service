package pdf_test

import (
	"log"
	"os"
	"testing"

	"projects/doc/doc_service/pkg/transport"
	"projects/doc/doc_service/pkg/transport/connect"
	"projects/doc/doc_service/pkg/transport/formats/pdf"
)

func TestPdf_Perform(t *testing.T) {
	tr, err := transport.NewTransport("127.0.0.1:8030", "5fb24be9-084a-4b34-80e5-28f21595bae6", &connect.TCreate{
		Name:    "Test PDF",
		Comment: "Тестовый запуск формирования PDF",
	})
	if err != nil {
		log.Println("создание транспорта", err)
	}

	doc1, err := os.ReadFile("1.docx")
	if err != nil {
		log.Println("Загрузка документа", err)
	}
	// doc2, err := os.ReadFile("2.docx")
	// if err != nil {
	// 	log.Println("Загрузка документа", err)
	// }

	res, err := tr.PdfPerform(pdf.TFiles{
		Files: []*pdf.TFile{
			{
				FileData: doc1,
				Ext:      "docx",
				Name:     "1",
			},
			// {
			// 	FileData: doc2,
			// 	Ext:      "docx",
			// 	Name:     "test_2",
			// 	Params: pdf.TParams{
			// 		// Join:     true,
			// 		Rotation: true,
			// 	},
			// },
		},
	})

	if err != nil {
		log.Println("Отправка данных", err)
	}

	if res == nil {
		return
	}

	f, err := os.Create("test." + res.Ext)
	if err != nil {
		log.Println("создание файла", err)
	}

	f.Write(res.FileData)

	f.Close()

	log.Println(res.Ext)
}
