package render

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func (t *TRenderDocOne) download() error {
	// HTTP endpoint
	postUrl := fmt.Sprintf("%v?DocumentFormat=%v", t.link, t.format)

	r, err := http.NewRequest("GET", postUrl, bytes.NewBuffer(nil))
	if err != nil {
		return fmt.Errorf("TRenderDocOne.download(): формирования запроса, err=%w", err)
	}

	r.Header.Add("Authorization", fmt.Sprintf("Basic %v", t.auth.GetAuthBase64()))

	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("TRenderDocOne.download(): выполнение запроса, err=%w", err)
	}

	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("TRenderDocOne.download(): чтение файла, err=%w", err)
	}

	t.fileDocOne.Write(resp)
	t.fileExtDocOne = t.format

	return nil
}
