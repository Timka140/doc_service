package fill_xlsx

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

type tSelectResource struct {
	File     string `json:"file"`
	Comment  string `json:"comment"`
	FuncEdit string `json:"func_edit"`
}

type tSelectTemplate struct {
	Type     string                     `json:"type"`
	Resource map[string]tSelectResource `json:"resource"`
}

// Выбор нужного шаблона по коду документа
func (t *TFillXlsx) select_template(code string) (*excelize.File, string, error) {
	lBytes, err := os.ReadFile(filepath.Join(t.templates_path, "/xlsx/list.json"))
	if err != nil {
		return nil, "", fmt.Errorf("TFillXlsx.select_template(): не удалось загрузить список файлов, err=%w", err)
	}

	var list tSelectTemplate
	err = json.Unmarshal(lBytes, &list)
	if err != nil {
		return nil, "", fmt.Errorf("TFillXlsx.select_template(): ошибка структуры файла list.json, err=%w", err)
	}

	path, ok := list.Resource[code]
	if !ok {
		return nil, "", fmt.Errorf("TFillXlsx.select_template(): в файле list.json не найден, code=%v", code)
	}

	rXlsx, err := excelize.OpenFile(filepath.Join(t.templates_path, "/xlsx/", path.File))
	if err != nil {
		text := fmt.Sprintf("TFillXlsx.select_template(): Code: %v, error: %v", code, err)
		log.Println(text)
		return nil, "", fmt.Errorf(text)
	}

	return rXlsx, path.FuncEdit, nil
}
