package xlsx

import (
	"encoding/json"
	"fmt"
	"os"

	"projects/doc/doc_service/pkg/cons"
	"projects/doc/doc_service/pkg/transport/methods"
)

// Возвращает структуру запакованного документа документа
func (t *tFillXlsx) GetDocument() (file methods.TFile) {
	file = methods.TFile{
		FileData: t.fileXlsx.Bytes(),
		Ext:      cons.CExtXlsx,
		Name:     t.params.NameFile,
	}

	return file
}

// Возвращает файл docx в байтовом виде
func (t *tFillXlsx) ReadBytes() (data []byte, err error) {

	pack := methods.TGenerateReportRespPack{
		Files: []*methods.TFile{
			{
				FileData: t.fileXlsx.Bytes(),
				Ext:      cons.CExtXlsx,
				Name:     t.params.NameFile,
			},
		},
	}

	data, err = json.Marshal(pack)
	if err != nil {
		return nil, fmt.Errorf("TFillXlsx.ReadBytes(): закрытие файла docx, err=%w", err)
	}

	return data, nil
}

// Записывает файл docx по нужному пути
func (t *tFillXlsx) WriteToPath(path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("TFillXlsx.WriteToPath(): создание файла docx, err=%w", err)
	}

	f.Write(t.fileXlsx.Bytes())

	err = f.Close()
	if err != nil {
		return fmt.Errorf("TFillXlsx.WriteToPath(): запись файла docx, err=%w", err)
	}
	return nil
}
