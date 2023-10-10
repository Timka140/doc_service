package methods_pdf

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"projects/doc/doc_service/pkg/transport/methods"

	"github.com/google/uuid"
)

func (t *TMethodsPDF) MargePDF(in *methods.TGenerateReportGroup) error {
	if len(t.marge_files) == 0 {
		return nil
	}

	if runtime.GOOS == "windows" {
		return nil
	}

	name := fmt.Sprintf("%v.%v", uuid.NewString(), t.doc.File.Ext)

	fPath := filepath.Join(t.catalog_path, name)
	t.marge_files = append(t.marge_files, fPath) // указываю название файла

	cmd := exec.Command("pdfunite", t.marge_files...)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		text := fmt.Sprintf("[ERROR] MargePDFCmd Run console,\n\terror: %v", err)
		log.Println(text)
		return fmt.Errorf(text)
	}

	fPDF, err := os.ReadFile(fPath)
	if err != nil {
		return fmt.Errorf("methods_pdf.MargePDF() : чтение файла PDF, err=%w", err)
	}

	//Добавляю новый файл
	in.ReportFiles = append(in.ReportFiles, &methods.TReport{
		File: methods.TFile{
			Ext:      ExtPDF,
			FileData: fPDF,
			Name:     "marge",
		},
	})

	return nil
}

func (t *TMethodsPDF) AddMargeList() {
	if !t.doc.Params.Join {
		return
	}
	t.marge_files = append(t.marge_files, t.file_path)
}

func (t *TMethodsPDF) RemoveMargeFile(in []*methods.TReport) (out []*methods.TReport) {
	buf := []*methods.TReport{}

	for _, val := range in {
		if val.Params.Join {
			continue
		}
		buf = append(buf, val)
	}

	return buf
}
