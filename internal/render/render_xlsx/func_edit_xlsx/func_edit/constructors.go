package func_edit_xlsx

import (
	"fmt"
	"log"
)

var (
	funcEditXlsx IFuncEditXlsx
)

func GetConstructorXlsx() (IFuncEditXlsx, error) {
	if funcEditXlsx == nil {
		return nil, fmt.Errorf("func_edit.GetConstructorXlsx(): функции редактирования не заданы")
	}

	return funcEditXlsx, nil
}

func init() {
	var err error
	funcEditXlsx, err = NewFuncEditXlsx()
	if err != nil {
		log.Printf("func_edit.init(): инициализация функций редактирования xlsx, err=%v", err)
	}
}
