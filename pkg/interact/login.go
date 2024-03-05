package interact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (t *TInteract) Login(login, password string) error {
	data := map[string]string{
		"login":    login,
		"password": password,
	}

	// Кодируем данные в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("Login(): формирования json, err=%w", err)
	}

	// Создаем HTTP-запрос типа POST с телом в формате JSON
	req, err := http.NewRequest("POST", fmt.Sprintf("%v://%v/api/login", t.Protocol, t.Address), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("Login(): создание запроса, err=%w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос и получаем ответ
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Login(): отправка запроса, err=%w", err)
	}
	defer resp.Body.Close()

	// Выводим статус ответа
	switch resp.StatusCode {
	case http.StatusOK:
	default:
		return fmt.Errorf("Login(): ответ, code=%v %v", resp.StatusCode, resp.Status)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Login(): чтение ответа, err=%w", err)
	}

	res := map[string]interface{}{}
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return fmt.Errorf("Login(): чтение ответа, err=%w", err)
	}

	if token, ok := res["token"].(string); ok {
		t.token = token
	}

	return nil
}
