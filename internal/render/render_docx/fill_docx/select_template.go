package fill_docx

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type tSelectResource struct {
	File    string `json:"file"`
	Comment string `json:"comment"`
}

type tSelectTemplate struct {
	Type     string                     `json:"type"`
	Resource map[string]tSelectResource `json:"resource"`
}

// Выбор нужного шаблона по коду документа
func (t *TFillDocx) select_template(code string) ([]byte, error) {
	lBytes, err := os.ReadFile(filepath.Join(t.templates_path, "/docx/list.json"))
	if err != nil {
		return nil, fmt.Errorf("TFillDocx.select_template(): не удалось загрузить список файлов, err=%w", err)
	}

	var list tSelectTemplate
	err = json.Unmarshal(lBytes, &list)
	if err != nil {
		return nil, fmt.Errorf("TFillDocx.select_template(): ошибка структуры файла list.json, err=%w", err)
	}

	path, ok := list.Resource[code]
	if !ok {
		return nil, fmt.Errorf("TFillDocx.select_template(): в файле list.json не найден, code=%v", code)
	}

	file, err := os.ReadFile(filepath.Join(t.templates_path, "/docx/", path.File))
	if err != nil {
		return nil, fmt.Errorf("TFillDocx.select_template(): чтение шаблона, code=%v", code)
	}

	return file, nil
}
