package docx_test

import (
	"log"
	"os"
	"testing"

	"projects/doc/doc_service/pkg/transport"
	"projects/doc/doc_service/pkg/transport/methods"
)

func TestTDocx_Perform(t *testing.T) {
	tr, err := transport.NewTransport("127.0.0.1:8030")
	if err != nil {
		log.Println("создание транспорта", err)
	}

	res, err := tr.DocOnePerform("b72cedea-61b6-11ed-9e58-02c16399bff6", methods.TParams{Rotation: false, NameFile: "doc_one_test"}, map[string]interface{}{
		"Формальное название потребителя": []string{"Тестовые данные"},
		"номер договора":                  []string{"24352345324523454"},
		"Адрес филиала":                   []string{"24352345324523454"},
		"filial_name":                     []string{"filial тестовые данные"},
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
