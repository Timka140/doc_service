package fill_pdf

import (
	"encoding/json"
	"fmt"
	"os"

	"projects/doc/doc_service/pkg/cons"
	"projects/doc/doc_service/pkg/transport/methods"
)

// Возвращает структуру запакованного документа документа
func (t *TFillPdf) GetDocument() (file methods.TFile) {
	file = methods.TFile{
		FileData: t.filePdf.Bytes(),
		Ext:      cons.CExtPdf,
		Name:     t.params.NameFile,
	}

	return file
}

// Возвращает файл docx в байтовом виде
func (t *TFillPdf) ReadBytes() (data []byte, err error) {

	pack := methods.TGenerateReportRespPack{
		Files: []*methods.TFile{
			{
				FileData: t.filePdf.Bytes(),
				Ext:      cons.CExtPdf,
				Name:     t.params.NameFile,
			},
		},
	}

	data, err = json.Marshal(pack)
	if err != nil {
		return nil, fmt.Errorf("fill_docx.ReadBytes(): закрытие файла docx, err=%w", err)
	}

	return data, nil
}

// Записывает файл docx по нужному пути
func (t *TFillPdf) WriteToPath(path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("TFillDocx.WriteToPath(): создание файла docx, err=%w", err)
	}

	f.Write(t.filePdf.Bytes())

	err = f.Close()
	if err != nil {
		return fmt.Errorf("TFillDocx.WriteToPath(): запись файла docx, err=%w", err)
	}
	return nil
}
