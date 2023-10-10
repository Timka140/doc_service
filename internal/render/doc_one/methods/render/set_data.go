package render

import (
	"encoding/json"
	"fmt"
)

type TDocOneData struct {
	Data   map[string][]string `json:"params"`
	Format string              `json:"format"`
	Code   string              `json:"code"`
}

/*
//Пример входящих данных для калоризатора

	in = map[string][]string{
		"Формальное название потребителя": {"Тестовые данные"},
		"номер договора":                  {"24352345324523454"},
		"Адрес филиала":                   {"24352345324523454"},
		"filial_name":                     {"filial тестовые данные"},
	}
*/
func (t *TRenderDocOne) setData(in *map[string][]string) error {

	body, err := json.Marshal(map[string]interface{}{
		"Data": in,
	})
	if err != nil {
		return fmt.Errorf("SetData(): формирование данных, err=%w", err)
	}

	t.body = body
	return nil
}

func (t *TRenderDocOne) setTemplateID(in string) error {
	if in == "" {
		return fmt.Errorf("SetData(): TemplateID не задан")
	}
	t.templateID = in
	return nil
}
