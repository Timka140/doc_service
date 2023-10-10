package constructors_xlsx

import (
	"bytes"
	"fmt"
	"log"

	formats_xlsx "projects/doc/doc_service/pkg/transport/formats/xlsx"

	"github.com/xuri/excelize/v2"
)

type TClaim struct {
	fXlsx  *excelize.File
	values *formats_xlsx.TValuesRender
	file   *bytes.Buffer
}

func NewClaim(in *TDependenciesRender) IConstructorsXlsx {
	t := &TClaim{
		fXlsx: in.FXlsx,
	}

	return t
}

func (t *TClaim) SetData(in *formats_xlsx.TValuesRender) error {
	if in == nil {
		return fmt.Errorf("TClaim.SetData(): данные пустые")
	}
	t.values = in
	return nil
}

func (t *TClaim) Render() error {
	var err error

	// Была 11 строка
	startRow := 10
	for i := 0; i < len(t.values.DataTable)-1; i++ {
		row := startRow + i
		err = t.fXlsx.DuplicateRow("Лист1", row-1)
		if err != nil {
			return fmt.Errorf("TClaim.Render(): копирование строки, err=%w", err)
		}

		for k, v := range t.values.DataTable[i] {
			err = t.fXlsx.SetCellValue("Лист1", fmt.Sprintf("%v%v", k, row), v)
			if err != nil {
				return fmt.Errorf("TClaim.Render(): установка значения в ячейку, err=%w", err)
			}
		}
	}

	//Удаляю строку форматирования
	err = t.fXlsx.RemoveRow("Лист1", startRow-1)
	if err != nil {
		return fmt.Errorf("TClaim.Render(): удаление строки, err=%w", err)
	}

	posRow := startRow + len(t.values.DataTable) - 1

	// Итого
	for k, v := range t.values.DataTable[len(t.values.DataTable)-1] {
		if k == "A" {
			continue
		}
		err = t.fXlsx.SetCellValue("Лист1", fmt.Sprintf("%v%v", k, posRow-1), v)
		if err != nil {
			return fmt.Errorf("TClaim.Render(): установка строки, err=%w", err)
		}
	}

	// TODO return normal params render
	// ==============================================================
	A2 := fmt.Sprintf("Приложение к претензии № %v от %v",
		t.values.Data["claim_number"], t.values.Data["claim_date"])
	err = t.fXlsx.SetCellValue("Лист1", "A2", A2)
	if err != nil {
		return fmt.Errorf("TClaim.Render(): установка строки, err=%w", err)
	}

	A3 := fmt.Sprintf("%v (договор № %v от %v)",
		t.values.Data["partner_name"], t.values.Data["con_number"], t.values.Data["contract_date"])
	err = t.fXlsx.SetCellValue("Лист1", "A3", A3)
	if err != nil {
		return fmt.Errorf("TClaim.Render(): установка строки, err=%w", err)
	}

	// A4 := fmt.Sprintf("Период: %v - %v (по состоянию на %v)",
	// 	params["period_from"], params["period_to"], params["claim_date"])
	A4 := fmt.Sprintf("Период: %v",
		t.values.Data["period_from"])
	err = t.fXlsx.SetCellValue("Лист1", "A4", A4)
	if err != nil {
		return fmt.Errorf("TClaim.Render(): установка строки, err=%w", err)
	}

	// image := fmt.Sprintf("%v", t.values.Data["facsimile"])
	// opts := &excelize.GraphicOptions{
	// 	AltText:         "",
	// 	PrintObject:     nil,
	// 	Locked:          nil,
	// 	LockAspectRatio: false,
	// 	AutoFit:         false,
	// 	OffsetX:         0,
	// 	OffsetY:         0,
	// 	ScaleX:          0.25,
	// 	ScaleY:          0.25,
	// 	Hyperlink:       "",
	// 	HyperlinkType:   "",
	// 	Positioning:     "",
	// }
	// err = t.fXlsx.AddPicture("Лист1", fmt.Sprintf("B%v", posRow), image, opts)
	// if err != nil {
	// 	return fmt.Errorf("TClaim.Render(): установка строки, err=%w", err)
	// }

	C12 := fmt.Sprintf("%v", t.values.Data["manager_name"])
	err = t.fXlsx.SetCellValue("Лист1", fmt.Sprintf("C%v", posRow+2), C12)
	if err != nil {
		return fmt.Errorf("TClaim.Render(): установка строки, err=%w", err)
	}

	A15 := fmt.Sprintf("%v", t.values.Data["curator_name"])
	err = t.fXlsx.SetCellValue("Лист1", fmt.Sprintf("B%v", posRow+5), A15)
	if err != nil {
		return fmt.Errorf("TClaim.Render(): установка строки, err=%w", err)
	}

	A16 := fmt.Sprintf("Тел.: %v", t.values.Data["curator_phone"])
	err = t.fXlsx.SetCellValue("Лист1", fmt.Sprintf("B%v", posRow+6), A16)
	if err != nil {
		return fmt.Errorf("TClaim.Render(): установка строки, err=%w", err)
	}

	A17 := fmt.Sprintf("Сформировано автоматически %v в программном комплексе \"РАПИРА\"",
		t.values.Data["current_time"])
	err = t.fXlsx.SetCellValue("Лист1", fmt.Sprintf("A%v", posRow+7), A17)
	if err != nil {
		return fmt.Errorf("TClaim.Render(): установка строки, err=%w", err)
	}

	t.file, err = t.fXlsx.WriteToBuffer()
	if err != nil {
		return fmt.Errorf("TClaim.Render(): сохранение в буфер, err=%w", err)
	}

	return nil
}

func (t *TClaim) GetFileXlsx() (*bytes.Buffer, error) {
	return t.file, nil
}

func init() {
	err := constructors.Add("Claim", NewClaim)
	if err != nil {
		log.Printf("NewClaim(): не удалось добавить в конструктор")
	}
}
