package docx_test

import (
	"fmt"
	"log"
	"sync"
	"testing"

	"projects/doc/doc_service/pkg/transport"
	"projects/doc/doc_service/pkg/transport/methods"
)

func TestTDocx_Perform(t *testing.T) {
	var wg sync.WaitGroup

	for index := 0; index < 8; index++ {
		testWork(&wg, fmt.Sprintf("work_%v", index))
	}
	wg.Wait()
}

func testWork(wg *sync.WaitGroup, name string) {
	wg.Add(1)
	go func(wg *sync.WaitGroup, name string) {
		tr, err := transport.NewTransport("127.0.0.1:8030")
		if err != nil {
			log.Println("создание транспорта", err)
		}

		for index := 0; index < 10_000; index++ {

			res, err := tr.DocxPerform("1", methods.TParams{NameFile: fmt.Sprintf("test_%v_%v", name, index), ConvertPDF: false, Rotation: false},
				map[string]interface{}{
					"col_labels": []string{"fruit", "vegetable", "stone", "thing"},
					"tbl_contents": []interface{}{
						map[string]interface{}{"label": "yellow", "cols": []string{"banana", "capsicum", "pyrite", "taxi"}},
						map[string]interface{}{"label": "red", "cols": []string{"apple", "tomato", "cinnabar", "doubledecker"}},
						map[string]interface{}{"label": "green", "cols": []string{"guava", "cucumber", "aventurine", "card"}},
					},
				})

			if err != nil {
				log.Println("Отправка данных", err)
			}

			if res == nil {
				return
			}

			fmt.Println(res.Name)

			// f, err := os.Create(filepath.Join("res", res.Name+"."+res.Ext))
			// if err != nil {
			// 	log.Println("создание файла", err)
			// }

			// f.Write(res.FileData)

			// f.Close()

		}
		wg.Done()
	}(wg, name)
}
