package methods_get

import (
	"log"

	"github.com/gin-gonic/gin"
)

type TStartProcess struct {
	// ctx *gin.Context
}

func newStartProcess(in *TInGetPage) IGetPage {
	t := &TStartProcess{}

	return t
}

func (t *TStartProcess) GetPath() string {
	return "/start_process"
}

func (t *TStartProcess) GetContext(c *gin.Context) {

}

func init() {
	err := constructors.Add("StartProcess", newStartProcess)
	if err != nil {
		log.Printf("StartProcess(): не удалось добавить в конструктор")
	}
}
