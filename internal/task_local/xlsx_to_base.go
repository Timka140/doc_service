package task_local

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

type tInXlsxToBase struct {
	File bytes.Buffer
	DB   *sql.DB
	List string
}

func (t *TTaskLocal) xlsxToBase(in *tInXlsxToBase) ([]map[string]interface{}, error) {
	var err error
	table := make([]map[string]interface{}, 0)

	if in.File.Len() == 0 {
		return table, nil
	}

	var rXlsx *excelize.File

	rXlsx, err = excelize.OpenReader(&in.File)
	if err != nil {
		return nil, fmt.Errorf("NewXlsxToBase(): чтение файла xlsx, err=%w", err)
	}

	rows, err := rXlsx.Rows(in.List)
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
			key := header[iCel]
			line[key] = colCell
		}
		table = append(table, line)
		ind += 1
	}
	if err = rows.Close(); err != nil {
		return nil, fmt.Errorf("NewXlsxToBase(): закрытие файла xlsx, err=%w", err)
	}

	return table, nil
}
