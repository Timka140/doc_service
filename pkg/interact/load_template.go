package interact

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
)

func (t *TInteract) LoadTemplate(id string, file *bytes.Buffer) error {
	if id == "" {
		return fmt.Errorf("LoadTemplate(): id шаблона не задан")
	}

	if file == nil {
		return fmt.Errorf("LoadTemplate(): файл не задан")
	}

	// Создаем буфер для формирования multipart/form-data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Добавляем файл
	part, err := writer.CreateFormFile("file", "example.txt")
	if err != nil {
		return fmt.Errorf("LoadTemplate(): создание файла, err=%w", err)
	}

	if _, err := file.WriteTo(part); err != nil {
		return fmt.Errorf("LoadTemplate(): копирование файла в форму, err=%w", err)
	}

	// Добавляем параметр template_id
	if err := writer.WriteField("template_id", id); err != nil {
		return fmt.Errorf("LoadTemplate(): копирование файла в форму, err=%w", err)
	}

	// Закрываем writer и получаем содержимое формы
	if err := writer.Close(); err != nil {
		return fmt.Errorf("LoadTemplate(): закрытие writer, err=%w", err)
	}

	// Создаем HTTP-запрос типа POST с телом в формате JSON
	req, err := http.NewRequest("POST", fmt.Sprintf("%v://%v/api/load_template", t.Protocol, t.Address), body)
	if err != nil {
		return fmt.Errorf("LoadTemplate(): создание запроса, err=%w", err)
	}
	req.Header.Set("AccessToken", t.token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Отправляем запрос и получаем ответ
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("LoadTemplate(): отправка запроса, err=%w", err)
	}
	defer resp.Body.Close()

	// Выводим статус ответа
	switch resp.StatusCode {
	case http.StatusOK:
	default:
		return fmt.Errorf("LoadTemplate(): ответ, code=%v %v", resp.StatusCode, resp.Status)
	}
	return nil
}
