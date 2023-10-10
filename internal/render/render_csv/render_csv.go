package render_csv

// Заготовка под формирования Csv файлов
type IRenderCsv interface {
}

type TRenderCsv struct {
}

func NewRenderCsv() (IRenderCsv, error) {
	t := &TRenderCsv{}

	return t, nil
}
