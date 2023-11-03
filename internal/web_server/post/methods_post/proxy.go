package methods_post

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TProxy struct {
	// ctx *gin.Context
}

func newProxy(in *TInPostPage) IPostPage {
	t := &TProxy{}

	return t
}

func (t *TProxy) GetPath() string {
	return "/proxy"
}

func (t *TProxy) GetContext(c *gin.Context) {
	data := bytes.NewBuffer(nil)
	io.Copy(data, c.Request.Body)
	if data.Len() == 0 {
		return
	}
	params := make(map[string]interface{})
	err := json.Unmarshal(data.Bytes(), &params)
	// err := c.Request.ParseForm()
	if err != nil {
		log.Println("TLoginPost.GetContext(): чтение параметров, err=%w", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"state": "ok",
	})
}

func init() {
	err := constructors.Add("Proxy", newProxy)
	if err != nil {
		log.Printf("Proxy(): не удалось добавить в конструктор")
	}
}
