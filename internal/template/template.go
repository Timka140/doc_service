package template

import (
	"fmt"
	"os"
	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/pkg/types"
	"strconv"
	"time"
)

type tTemplate struct {
	templateID int64
	file       *types.TFile
}

// New -- принимает файл или id_template
func New(templateID interface{}) (types.ITemplate, error) {
	t := &tTemplate{}

	var err error
	switch v := templateID.(type) {
	case int64:
		if v == 0 {
			return nil, fmt.Errorf("template.New(): templateID шаблона не задан")
		}
		t.templateID = v
	case string:
		if v == "" {
			return nil, fmt.Errorf("template.New(): templateID шаблона не задан")
		}

		t.templateID, err = strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("template.New(): templateID не число, templateID = %v", v)
		}

	case nil:
		return nil, fmt.Errorf("template.New(): templateID не задан")
	default:
		return nil, fmt.Errorf("template.New(): неизвестный тип templateID, %v", v)
	}

	t.file, err = t.BaseLoad()
	if err != nil {
		return nil, fmt.Errorf("template.New(): загрузка, %v", err)
	}

	return t, nil
}

func (t *tTemplate) ToBase() error {
	pack, err := Pack(t.file)
	if err != nil {
		return fmt.Errorf("tTemplate.ToBase(): упаковка, err=%w", err)
	}

	err = db.DB.Table("templates").Where("id = ?", t.templateID).Update("data", pack).Error
	if err != nil {
		return fmt.Errorf("tTemplate.ToBase(): не удалось загрузить файл, err=%w", err)
	}
	return nil
}

func (t *tTemplate) BaseLoad() (*types.TFile, error) {
	var templates db.Templates
	err := db.DB.Table("templates").Select("data").Where("id = ?", t.templateID).Scan(&templates).Error
	if err != nil {
		return nil, fmt.Errorf("tTemplate.BaseLoad(): не удалось загрузить файл, err=%w", err)
	}

	//Возвращаю пустой файл
	if templates.Data == nil {
		return &types.TFile{}, nil
	}

	file, err := UnPack(templates.Data)
	if err != nil {
		return nil, fmt.Errorf("tTemplate.BaseLoad():распаковка, err=%w", err)
	}

	return file, nil
}

func (t *tTemplate) IsFile() bool {
	if _, err := os.Stat(t.file.PathTemplate); err == nil {
		return true
	}
	return false
}

// Template -- возвращает шаблон
func (t *tTemplate) Template() ([]byte, error) {
	data, err := os.ReadFile(t.file.PathTemplate)
	if err != nil {
		return nil, fmt.Errorf("tTemplate.BaseLoad(): чтение шаблона, err=%w", err)
	}
	return data, nil
}

func (t *tTemplate) Name() string {
	return t.file.Name
}

func (t *tTemplate) UpdateTime() string {
	tm, _ := time.Parse(time.RFC3339Nano, t.file.Update)

	return tm.Format("02.01.2006 15:04:05")
}
