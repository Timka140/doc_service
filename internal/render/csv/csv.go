package csv

// Заготовка под формирования Csv файлов
type IRenderCsv interface {
}

type tRenderCsv struct {
}

func NewRenderCsv() (IRenderCsv, error) {
	t := &tRenderCsv{}

	return t, nil
}
