package methods_pdf

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	ExtPDF = "pdf"
)

// Convert - конвертирует в формат .pdf любой файл понимаемый LibreOffice
// в том числе .xlsx, .docx, .txt, .fodt, .fods
func (t *TMethodsPDF) Convert() error {

	if !t.doc.Params.ConvertPDF && !t.doc.Params.Join {
		return nil
	}

	// tm := time.Now()
	soffice_name := "soffice"
	if runtime.GOOS == "windows" {
		soffice_name = "soffice.exe"
	}
	cmd := exec.Command(soffice_name, "--headless", "--convert-to", "pdf", "--outdir", t.catalog_path, t.file_path)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Run()
	// duration := time.Since(tm)
	// log.Println("ConvertToPDF(): End convert, time = ", duration)
	if err != nil {
		return fmt.Errorf("methods_pdf.Convert(), Run: %v,\n\terror: %w", t.file_path, err)
	}

	pdf_path := strings.Replace(t.file_path, t.doc.File.Ext, "pdf", -1)

	fPDF, err := os.ReadFile(pdf_path)
	if err != nil {
		return fmt.Errorf("methods_pdf.Convert() : чтение файла PDF, err=%w", err)
	}

	t.file_path = pdf_path
	t.doc.File.FileData = fPDF
	t.doc.File.Ext = ExtPDF
	return nil
}
