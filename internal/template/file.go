package template

import (
	"encoding/json"
	"fmt"
	"projects/doc/doc_service/pkg/types"
)

// Pack -- упаковка файла
func Pack(file *types.TFile) ([]byte, error) {
	pack, err := json.Marshal(file)
	if err != nil {
		return nil, fmt.Errorf("template.Pack(): не удалось упаковать файл, err=%w", err)
	}
	return pack, nil
}

// UnPack -- распаковка файла
func UnPack(bin []byte) (*types.TFile, error) {
	file := &types.TFile{}
	err := json.Unmarshal(bin, file)
	if err != nil {
		return nil, fmt.Errorf("template.Pack(): не удалось распаковать файл, err=%w", err)
	}
	return file, nil
}
