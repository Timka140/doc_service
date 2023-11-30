package docx_test

// import (
// 	"log"
// 	"os"
// 	"testing"

// 	"projects/doc/doc_service/pkg/transport"
// 	"projects/doc/doc_service/pkg/transport/methods"
// )

// // func TestTXlsx_Perform(t *testing.T) {
// // 	tr, err := transport.NewTransport("127.0.0.1:8030")
// // 	if err != nil {
// // 		log.Println("создание транспорта", err)
// // 	}

// // 	res, err := tr.XlsxPerform("1", methods.TParams{
// // 		NameFile:   "test_claim",
// // 		ConvertPDF: true,
// // 		// Rotation:   true,
// // 	},
// // 		map[string]interface{}{
// // 			"col_labels": []string{"fruit", "vegetable", "stone", "thing"},
// // 			"tbl_contents": []interface{}{
// // 				map[string]interface{}{"label": "yellow", "cols": []string{"banana", "capsicum", "pyrite", "taxi"}},
// // 				map[string]interface{}{"label": "red", "cols": []string{"apple", "tomato", "cinnabar", "doubledecker"}},
// // 				map[string]interface{}{"label": "green", "cols": []string{"guava", "cucumber", "aventurine", "card"}},
// // 			},
// // 		})

// // 	if err != nil {
// // 		log.Println("Отправка данных", err)
// // 	}

// // 	if res == nil {
// // 		return
// // 	}

// // 	f, err := os.Create(res.Name + "." + res.Ext)
// // 	if err != nil {
// // 		log.Println("создание файла", err)
// // 	}

// // 	f.Write(res.FileData)

// // 	f.Close()

// // 	log.Println(res.Ext)
// // }
