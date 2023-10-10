package task_local

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/xuri/excelize/v2"
)

type TInXlsxToBase struct {
	File bytes.Buffer
	List string
}
type TXlsxTable struct {
	Headers []string
	Data    []map[string]interface{}
}
type TXlsxTables map[string]TXlsxTable

func XlsxToBase(in *TInXlsxToBase) (TXlsxTables, error) {
	var err error
	tables := make(TXlsxTables, 0)

	if in.File.Len() == 0 {
		return tables, nil
	}

	var rXlsx *excelize.File

	rXlsx, err = excelize.OpenReader(&in.File)
	if err != nil {
		return nil, fmt.Errorf("NewXlsxToBase(): чтение файла xlsx, err=%w", err)
	}

	sheets := rXlsx.GetSheetList()

	for _, sheet := range sheets {
		if !IsEngByLoop(sheet) {
			continue
		}
		table := TXlsxTable{
			Data: make([]map[string]interface{}, 0),
		}
		rows, err := rXlsx.Rows(sheet)
		if err != nil {
			return nil, fmt.Errorf("NewXlsxToBase(): чтение файла xlsx, err=%w", err)
		}

		var header []string
		var ind int
		for rows.Next() {
			row, err := rows.Columns()
			if err != nil {
				log.Println(err)
			}
			if ind == 0 {
				header = row
				for ind, v := range header {
					header[ind] = strings.ToLower(v)
				}
				ind += 1
				continue
			}
			if ind == 1 {
				ind += 1
				continue
			}
			line := make(map[string]interface{})
			for iCel, colCell := range row {
				if iCel > len(header)-1 {
					break
				}
				key := header[iCel]
				line[key] = colCell
			}
			table.Data = append(table.Data, line)
			ind += 1
		}
		if err = rows.Close(); err != nil {
			return nil, fmt.Errorf("NewXlsxToBase(): закрытие файла xlsx, err=%w", err)
		}

		table.Headers = header

		tables[sheet] = table
	}

	return tables, nil
}

func IsEngByLoop(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
