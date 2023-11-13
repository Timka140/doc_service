package static

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func NewStatic(router *gin.Engine) (err error) {
	if router == nil {
		return fmt.Errorf("NewStatic(): роутер не установлен")
	}

	static := os.Getenv("Static")
	if static == "" {
		return fmt.Errorf("NewStatic(): отсутствует путь к статичным файлам")
	}

	router.StaticFile("/favicon.ico", filepath.Join(static, "favicon.ico"))
	// router.StaticFS("dist", http.Dir(filepath.Join(static)))

	router.StaticFS("/css", http.Dir(filepath.Join(static, "css")))
	router.StaticFS("/js", http.Dir(filepath.Join(static, "js")))
	router.StaticFS("/assets", http.Dir(filepath.Join(static, "assets")))

	return nil
}
