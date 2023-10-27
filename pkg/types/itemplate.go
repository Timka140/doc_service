package types

type TFile struct {
	Name   string `json:"name"`
	Update string `json:"update"`
	Data   []byte `json:"data"`
	Ext    string `json:"ext"`
}

// ITemplate -- работа с шаблонами
type ITemplate interface {
	// ToBase -- загружает файл в базу
	ToBase() error
	// BaseLoad -- загружает файл
	BaseLoad() (*TFile, error)
	// IsFile -- проверяет наличие файла
	IsFile() bool
	// Name -- имя файла
	Name() string
	// UpdateTime -- последнее обновление файла
	UpdateTime() string
}
