package docx_test

import (
	"log"
	"os"
	"testing"

	"projects/doc/doc_service/pkg/transport"
	"projects/doc/doc_service/pkg/transport/methods"
)

func TestTXlsx_Perform(t *testing.T) {
	tr, err := transport.NewTransport("127.0.0.1:8030")
	if err != nil {
		log.Println("создание транспорта", err)
	}

	res, err := tr.XlsxPerform("1", methods.TParams{
		NameFile:   "test_claim",
		ConvertPDF: true,
		// Rotation:   true,
	},
		map[string]interface{}{
			"Data": map[string]interface{}{
				"claim_number": "4444",
				"claim_date":   "13.11.1998",
				"partner_name": "ТЕСТ партнер ООО АтомЭнергоСбыт",
				"facsimile":    "",
			},
			"DataTable": []map[string]interface{}{
				{
					"A": "5110204838/086943 от 30.06.2023",
					"B": "93 336,82",
				},
			},
		})

	if err != nil {
		log.Println("Отправка данных", err)
	}

	if res == nil {
		return
	}

	f, err := os.Create(res.Name + "." + res.Ext)
	if err != nil {
		log.Println("создание файла", err)
	}

	f.Write(res.FileData)

	f.Close()

	log.Println(res.Ext)
}
