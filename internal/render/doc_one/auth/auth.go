package auth

import (
	"encoding/base64"
	"fmt"
	"os"
)

type TAuth struct {
	login    string
	password string
	url      string

	b64 string
}

type IAuth interface {
	GetAuthBase64() string
	GetURL() string
}

func NewAuth() (IAuth, error) {

	doc_one_url := os.Getenv("DOC_ONE")
	if doc_one_url == "" {
		return nil, fmt.Errorf("NewDocOne(): адрес DOC_ONE не установлен в .env")
	}

	login := os.Getenv("LOGIN_DOC_ONE")
	if login == "" {
		return nil, fmt.Errorf("NewAuth(): не задан логин для doc one")
	}
	password := os.Getenv("PASSWORD_DOC_ONE")
	if password == "" {
		return nil, fmt.Errorf("NewAuth(): не задан пароль для doc one")
	}

	t := &TAuth{
		login:    login,
		password: password,
		url:      doc_one_url,
		b64:      base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", login, password))),
	}

	return t, nil
}

func (t *TAuth) GetAuthBase64() string {
	return t.b64
}

func (t *TAuth) GetURL() string {
	return t.url
}
