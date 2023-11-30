package services

import (
	"log"
	"os"
)

var Services IServices

func init() {
	var err error
	Services, err = New()
	if err != nil {
		log.Printf("services.init(): создание списка сервисов инициаторов, err=%v", err)
		os.Exit(1)
	}
}
