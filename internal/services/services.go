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

func ServicesList() []map[string]interface{} {
	var data []map[string]interface{}
	Services.Range(func(srv IService) {
		data = append(data, map[string]interface{}{
			"name":    srv.Name(),
			"comment": srv.Comment(),
			"ping":    srv.Ping(),
		})
	})

	return data
}
