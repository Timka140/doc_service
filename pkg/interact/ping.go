package interact

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// /v1/isLogin
func (t *TInteract) ping() error {
	data := map[string]string{
		"api": "true",
	}

	// Кодируем данные в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("ping(): формирования json, err=%w", err)
	}

	// Создаем HTTP-запрос типа POST с телом в формате JSON
	req, err := http.NewRequest("POST", fmt.Sprintf("%v://%v/api/v1/isLogin", t.Protocol, t.Address), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("ping(): создание запроса, err=%w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("AccessToken", t.token)

	// Отправляем запрос и получаем ответ
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ping(): отправка запроса, err=%w", err)
	}
	defer resp.Body.Close()

	// Выводим статус ответа
	switch resp.StatusCode {
	case http.StatusOK:
	default:
		return fmt.Errorf("ping(): ответ, code=%v %v", resp.StatusCode, resp.Status)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ping(): чтение ответа, err=%w", err)
	}

	res := map[string]interface{}{}
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return fmt.Errorf("ping(): чтение ответа, err=%w", err)
	}

	if login, ok := res["login"].(bool); ok && !login {
		return fmt.Errorf("ping(): доступ закрыт")
	}

	return nil
}
