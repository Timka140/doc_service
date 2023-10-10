package methods_pdf

import (
	"fmt"
	"os"
	"path/filepath"
)

func (t *TMethodsPDF) CreateFile(name string) error {

	if name == "" {
		return fmt.Errorf("TMethodsPDF.SetPathFile(): имя не задано")
	}
	t.file_path = filepath.Join(t.catalog_path, name)

	f, err := os.Create(t.file_path)
	if err != nil {
		return fmt.Errorf("TMethodsPDF.CreateFile() : создание файла PDF, err=%w", err)
	}

	f.Write(t.doc.File.FileData)
	f.Close()
	return nil
}

func (t *TMethodsPDF) RemoveCatalog() error {
	err := os.RemoveAll(t.catalog_path)
	if err != nil {
		return fmt.Errorf("TMethodsPDF.CreateFile() : удаление файла PDF, err=%w", err)
	}
	return nil
}

func (t *TMethodsPDF) CrateCatalog(catalog string) error {
	if catalog == "" {
		return fmt.Errorf("TMethodsPDF.SetPathFile(): каталог не задан")
	}

	err := os.MkdirAll(filepath.Join(t.out_pdf, catalog), 0755)
	if err != nil {
		return fmt.Errorf("TMethodsPDF.SetPathFile(): создание папок, err=%w", err)
	}

	t.catalog_path = filepath.Join(t.out_pdf, catalog)

	return nil
}

func (t *TMethodsPDF) GetPathFile() string {
	return t.file_path
}
