package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// TemplateInit - Загрузка шаблонов
func NewTemplates() (temp *template.Template, err error) {
	frontSrc := os.Getenv("FrontSrc") // расположение шаблонов
	if frontSrc == "" {
		return nil, fmt.Errorf("TemplateInit(): env FRONT_SRC is empty")
	}
	temp, err = template.ParseGlob(filepath.Join(frontSrc, `*.html`))
	if err != nil {
		return nil, fmt.Errorf("NewTemplate(): загрузка шаблонов, err=%w", err)
	}

	return
}
