package types

type TFile struct {
	Name         string `json:"name"`
	Update       string `json:"update"`
	Ext          string `json:"ext"`
	PathTemplate string `json:"path_template"`
}

// ITemplate -- работа с шаблонами
type ITemplate interface {
	// ToBase -- загружает файл в базу
	ToBase() error
	// BaseLoad -- загружает файл
	BaseLoad() (*TFile, error)
	// IsFile -- проверяет наличие файла
	IsFile() bool
	// Template -- возвращает шаблон
	Template() ([]byte, error)

	// Name -- имя файла
	Name() string
	// UpdateTime -- последнее обновление файла
	UpdateTime() string
}
