package docx_test

import (
	"log"
	"os"
	"testing"

	"projects/doc/doc_service/pkg/transport"
	formats_xlsx "projects/doc/doc_service/pkg/transport/formats/xlsx"
	"projects/doc/doc_service/pkg/transport/methods"
)

func TestTDocx_Perform(t *testing.T) {
	tr, err := transport.NewTransport("127.0.0.1:8030")
	if err != nil {
		log.Println("создание транспорта", err)
	}

	group, err := tr.SendGroupFile()
	if err != nil {
		log.Println("инициализация открытия", err)
	}

	err = group.DocxPerform("1", methods.TParams{Rotation: false, NameFile: "claim", Join: true}, map[string]interface{}{
		"col_labels": []string{"fruit", "vegetable", "stone", "thing"},
		"tbl_contents": []interface{}{
			map[string]interface{}{"label": "yellow", "cols": []string{"banana", "capsicum", "pyrite", "taxi"}},
			map[string]interface{}{"label": "red", "cols": []string{"apple", "tomato", "cinnabar", "doubledecker"}},
			map[string]interface{}{"label": "green", "cols": []string{"guava", "cucumber", "aventurine", "card"}},
		},
	})
	if err != nil {
		log.Println("Добавление файла", err)
	}

	if err != nil {
		log.Println("Добавление файла", err)
	}

	err = group.XlsxPerform("1", methods.TParams{Rotation: true, NameFile: "claim_table", Join: true}, formats_xlsx.TValuesRender{
		Data: map[string]interface{}{
			"manager_name": "test",
		},
		DataTable: []map[string]interface{}{
			{
				"A": "5110204838/086943 от 30.06.2023",
				"B": "93 336,82",
			},
		},
	})
	if err != nil {
		log.Println("Добавление файла", err)
	}

	res, err := group.Send()
	if err != nil {
		log.Println("Отправка данных", err)
	}

	if res == nil {
		return
	}

	for _, val := range res {
		f, err := os.Create(val.Name + "." + val.Ext)
		if err != nil {
			log.Println("создание файла", err)
		}

		f.Write(val.FileData)

		f.Close()
	}
}
