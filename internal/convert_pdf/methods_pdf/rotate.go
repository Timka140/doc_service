package methods_pdf

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/google/uuid"
)

func (t *TMethodsPDF) Rotate() error {
	if !t.doc.Params.Rotation {
		return nil
	}

	if t.doc.File.Ext != ExtPDF {
		return nil
	}
	if runtime.GOOS == "windows" {
		return nil
	}

	// Аргументы команды pdftk

	in := t.file_path
	out := filepath.Join(t.catalog_path, uuid.NewString()) //fmt.Sprintf("%v/%v/rotate90_claim_%v.pdf", t.outFiles, in.Data.LawsuitId, in.Data.LawsuitId)

	args := []string{
		in,
		"cat",
		"1-endeast",
		"output",
		out,
	}

	pdftk_name := "pdftk"
	// if runtime.GOOS == "windows" {
	// 	pdftk_name = "pdftk.exe"
	// }

	cmd := exec.Command(pdftk_name, args...)

	var outB, errB bytes.Buffer
	cmd.Stdout = &outB
	cmd.Stderr = &errB
	err := cmd.Run()
	if err != nil {
		io.Copy(os.Stdout, &outB)
		io.Copy(os.Stdout, &errB)
		text := fmt.Sprintf("[ERROR] RotatePDF разворот PDF: %v", cmd.Err)
		// log.Println("Возможно не установлен пакет:")
		// log.Println("sudo apt install pdftk-java")
		return fmt.Errorf(text)
	}

	del_cmd := "rm"
	// if runtime.GOOS == "windows" {
	// 	del_cmd = "rd" //"rd"
	// }

	cmd = exec.Command(del_cmd, []string{
		in,
	}...)

	cmd.Stdout = &outB
	cmd.Stderr = &errB
	err = cmd.Run()
	if err != nil {
		io.Copy(os.Stdout, &outB)
		io.Copy(os.Stdout, &errB)
		text := fmt.Sprintf("[ERROR] RotatePDF удаление старого файла: %v", cmd.Err)
		// log.Println("Возможно не установлен пакет:")
		// log.Println("sudo apt install pdftk-java")
		return fmt.Errorf(text)
	}

	mv_cmd := "mv"
	if runtime.GOOS == "windows" {
		mv_cmd = "move" //"rd"
	}

	cmd = exec.Command(mv_cmd, []string{
		out,
		in,
	}...)

	cmd.Stdout = &outB
	cmd.Stderr = &errB
	err = cmd.Run()
	if err != nil {
		io.Copy(os.Stdout, &outB)
		io.Copy(os.Stdout, &errB)
		text := fmt.Sprintf("[ERROR] RotatePDF переименование файла: %v", cmd.Err)
		// log.Println("Возможно не установлен пакет:")
		// log.Println("sudo apt install pdftk-java")
		return fmt.Errorf(text)
	}

	fPDF, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("methods_pdf.Rotate() : чтение файла PDF, err=%w", err)
	}
	t.doc.File.FileData = fPDF

	return nil
}
