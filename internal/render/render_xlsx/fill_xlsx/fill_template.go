package fill_xlsx

import (
	"fmt"

	"projects/doc/doc_service/internal/render/render_xlsx/func_edit_xlsx/constructors_xlsx"
	formats_xlsx "projects/doc/doc_service/pkg/transport/formats/xlsx"
)

// Заполняет шаблон необходимыми данными
func (t *TFillXlsx) fill_template(funcEdit string, params map[string]interface{}) (err error) {

	fill := constructors_xlsx.NewConstructor()

	constructor, err := fill.Get(funcEdit)
	if err != nil {
		return fmt.Errorf("TFillXlsx.fill_template(): чтение конструктора")
	}

	fl := constructor(&constructors_xlsx.TDependenciesRender{
		FXlsx: t.fXlsx,
	})

	data, ok := params["Data"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("TFillXlsx.fill_template(): параметр Data неуказан")
	}

	table, ok := params["DataTable"].([]interface{})
	if !ok {
		return fmt.Errorf("TFillXlsx.fill_template(): параметр DataTable неуказан")
	}

	dataTable := make([]map[string]interface{}, len(table))

	for ind, itm := range table {
		v, ok := itm.(map[string]interface{})
		if !ok {
			continue
		}
		dataTable[ind] = v
	}

	err = fl.SetData(&formats_xlsx.TValuesRender{
		Data:      data,
		DataTable: dataTable,
	})
	if err != nil {
		return fmt.Errorf("TFillXlsx.fill_template(): установка данных, err=%w", err)
	}

	err = fl.Render()
	if err != nil {
		return fmt.Errorf("TFillXlsx.fill_template(): формирование документа, err=%w", err)
	}

	file, err := fl.GetFileXlsx()
	if err != nil {
		return fmt.Errorf("TFillXlsx.fill_template(): чтение документа, err=%w", err)
	}

	t.fileXlsx = file

	return nil
}
