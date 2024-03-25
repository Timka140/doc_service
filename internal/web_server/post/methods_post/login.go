package methods_post

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/gin-gonic/gin"
)

type TLoginPost struct {
	// ctx *gin.Context
}

func newLoginPost(in *TInPostPage) IPostPage {
	t := &TLoginPost{}

	return t
}

func (t *TLoginPost) GetPath() string {
	return "/v1/login"
}

func (t *TLoginPost) GetContext(c *gin.Context) {
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

	login, ok := params["login"].(string)
	if !ok {
		return
	}

	password, ok := params["password"].(string)
	if !ok {
		return
	}

	// login = "bondarenkotg"
	// password = "B~G|sUGP7bN%"

	// login := c.Request.FormValue("login")
	// password := c.Request.FormValue("password")

	key := md5.Sum([]byte(fmt.Sprintf("%v:docGenerator:%v", login, password)))
	hash := hex.EncodeToString(key[:])

	// hash := ""
	ses, err := sessions.NewSession(hash)
	if err != nil {
		log.Printf("GetContext(): авторизация, err=%v", err)
		return
	}

	if ses == nil {
		c.Status(http.StatusUnauthorized)
		c.Redirect(http.StatusTemporaryRedirect, "/gui/login")
		// redirect.Redirect(c, "/gui/login")
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  0,
			"message": "Неверный логин или пароль",
		})
		return
	}

	err = sessions.Ses.Add(ses.Token(), ses)
	if err != nil {
		log.Printf("GetContext(): запись сессии, err=%v", err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:  "AccessToken",
		Value: ses.Token(),
		Path:  "/",
		// MaxAge:   3600,
		// HttpOnly: true,
		// Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	c.Redirect(http.StatusTemporaryRedirect, "/gui/")
	resp := map[string]interface{}{
		"status":    1,
		"token":     ses.Token(),
		"rights":    ses.RightsVue(),
		"name":      "",
		"last_name": "",
		"message":   "",
	}
	c.JSON(http.StatusOK, resp)
}

func init() {
	err := constructors.Add("NewLoginPost", newLoginPost)
	if err != nil {
		log.Printf("NewLoginPost(): не удалось добавить в конструктор")
	}
}
