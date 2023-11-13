package docx_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"projects/doc/doc_service/pkg/transport"
	"projects/doc/doc_service/pkg/transport/formats/pdf"
)

func TestTDocx_Perform(t *testing.T) {
	tr, err := transport.NewTransport("127.0.0.1:8030")
	if err != nil {
		log.Println("создание транспорта", err)
	}

	res_docx, err := tr.DocxPerform("14",
		map[string]interface{}{
			"partner_full_name": "Заполнил новое поле",
			"test1":             "заполнил 2",
			"tbl_contents": []interface{}{
				map[string]interface{}{"label": "yellow", "cols": []string{"banana", "capsicum", "pyrite", "taxi"}},
				map[string]interface{}{"label": "red", "cols": []string{"apple", "tomato", "cinnabar", "doubledecker"}},
				map[string]interface{}{"label": "green", "cols": []string{"guava", "cucumber", "aventurine", "card"}},
			},
		})

	if err != nil {
		log.Println("создание файла", err)
	}

	res, err := tr.PdfPerform(pdf.TFiles{
		Files: []*pdf.TFile{
			{
				FileData: res_docx.FileData,
				Ext:      "docx",
				Name:     "test_1",
				Params: pdf.TParams{
					Rotation: true,
				},
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
		log.Println("создание файла", err)
	}

	f, err := os.Create(filepath.Join("res", "test."+res.Ext))
	if err != nil {
		log.Println("создание файла", err)
	}

	f.Write(res.FileData)

	f.Close()

}
