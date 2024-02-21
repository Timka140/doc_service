package methods_get

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TLiveness struct {
	// ctx *gin.Context
}

func newLiveness(in *TInGetPage) IGetPage {
	t := &TLiveness{}

	return t
}

func (t *TLiveness) GetPath() string {
	return "/api-doc/liveness"
}

func (t *TLiveness) GetContext(c *gin.Context) {
	c.Status(http.StatusOK)
}

func init() {
	err := constructors.Add("Liveness", newLiveness)
	if err != nil {
		log.Printf("Liveness(): не удалось добавить в конструктор")
	}
}
