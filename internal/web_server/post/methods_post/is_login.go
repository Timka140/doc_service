package methods_post

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/gin-gonic/gin"
)

type TIsLogin struct {
	// ctx *gin.Context
}

func newIsLogin(in *TInPostPage) IPostPage {
	t := &TIsLogin{}

	return t
}

func (t *TIsLogin) GetPath() string {
	return "/v1/isLogin"
}

func (t *TIsLogin) GetContext(c *gin.Context) {
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
	login := false
	ses := sessions.Ses.GetSes(c)
	if ses != nil {
		login = ses.Authorization()
	}

	path, ok := params["path"].(string)
	if ok && ses != nil {
		ses.SetCurrentPage(path)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"login": login,
	})
}

func init() {
	err := constructors.Add("IsLogin", newIsLogin)
	if err != nil {
		log.Printf("StartProcess(): не удалось добавить в конструктор")
	}
}
