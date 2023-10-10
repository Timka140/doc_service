package func_edit_xlsx

// Комплекс модулей для обработки файлов формата xlsx
type IFuncEditXlsx interface {
}

type TFuncEditXlsx struct {
}

func NewFuncEditXlsx() (IFuncEditXlsx, error) {
	t := &TFuncEditXlsx{}

	return t, nil
}
