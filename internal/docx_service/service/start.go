package service

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func (t *TService) startService() error {
	var err error
	// path, err := os.Getwd()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(path)
	if _, err := os.Stat("./docx_service"); os.IsNotExist(err) {
		return fmt.Errorf("start_service(): отсутствует микросервис для linux")
	}

	if _, err := os.Stat("./docx_service.exe"); os.IsNotExist(err) {
		return fmt.Errorf("start_service(): отсутствует микросервис для windows")
	}

	switch runtime.GOOS {
	case "windows":
		t.cmd = exec.Command("./docx_service.exe", t.pid, fmt.Sprintf("%v:%v", t.host, t.port), t.auth)
	case "linux":
		t.cmd = exec.Command("./docx_service", t.pid, fmt.Sprintf("%v:%v", t.host, t.port), t.auth)
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	t.cmd.Stdout = &stdout
	t.cmd.Stderr = &stderr
	err = t.cmd.Start()
	if err != nil {
		fmt.Println(stdout.String())
		fmt.Println(stderr.String())
		return fmt.Errorf("[ERROR] запуск сервиса docx_service: %v", err)
	}

	return nil
}

// terminateCommand() - завершение выполнения микросервиса
func (t *TService) terminateCommand() error {
	if t.cmd != nil && t.cmd.Process != nil {
		process := t.cmd.Process
		err := process.Kill()
		if err != nil {
			return fmt.Errorf("ошибка завершения процесса: %v", err)
		}
	}

	return nil
}

// listen() - ожидает команды из вне
func (t *TService) listen() {
	t.wg.Add(1)
	defer func() {
		t.wg.Done()
	}()

	close := <-t.cClose
	if !close.Close {
		return
	}
	err := t.terminateCommand()
	if err != nil {
		t.err = fmt.Errorf("NewStream(): остановка сервиса шаблонизации, err=%w", err)
		return
	}
	t.err = err
}
