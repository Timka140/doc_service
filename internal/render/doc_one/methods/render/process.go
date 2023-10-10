package render

import (
	"encoding/json"
	"fmt"
	"os"

	"projects/doc/doc_service/pkg/transport/methods"
)

func (t *TRenderDocOne) RenderDocOne(report *methods.TReport) error {
	var err error
	var data TDocOneData
	err = json.Unmarshal(report.Pack, &data)
	if err != nil {
		return fmt.Errorf("RenderDocOne(): распаковка пакета, err=%w", err)
	}

	t.params = report.Params

	err = t.setTemplateID(data.Code)
	if err != nil {
		return fmt.Errorf("RenderDocOne(): установка данных, err=%w", err)
	}

	err = t.setData(&data.Data)
	if err != nil {
		return fmt.Errorf("RenderDocOne(): установка данных, err=%w", err)
	}

	err = t.render()
	if err != nil {
		return fmt.Errorf("RenderDocOne(): рендер документа, err=%w", err)
	}

	err = t.download()
	if err != nil {
		return fmt.Errorf("RenderDocOne(): загрузка документа, err=%w", err)
	}

	return nil
}

func (t *TRenderDocOne) Stop() error {
	return nil
}

// Возвращает структуру запакованного документа документа
func (t *TRenderDocOne) GetDocument() (file methods.TFile) {
	file = methods.TFile{
		FileData: t.fileDocOne.Bytes(),
		Ext:      t.fileExtDocOne,
		Name:     t.params.NameFile,
	}

	return file
}

// Возвращает файл в байтовом виде
func (t *TRenderDocOne) ReadBytes() (data []byte, err error) {

	pack := methods.TGenerateReportRespPack{
		Files: []*methods.TFile{
			{
				FileData: t.fileDocOne.Bytes(),
				Ext:      t.fileExtDocOne,
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
func (t *TRenderDocOne) WriteToPath(path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("TFillDocx.WriteToPath(): создание файла docx, err=%w", err)
	}

	f.Write(t.fileDocOne.Bytes())

	err = f.Close()
	if err != nil {
		return fmt.Errorf("TFillDocx.WriteToPath(): запись файла docx, err=%w", err)
	}
	return nil
}
