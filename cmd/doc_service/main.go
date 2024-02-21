// package main -- пускач для обмена с сервисом DocOne
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/internal/grpc_server"
	"projects/doc/doc_service/internal/web_server"
)

func main() {
	serviceName := "doc_service"
	serviceID := "1"

	if len(os.Args) > 1 {
		serviceID = os.Args[1]
	}

	serviceHost := os.Getenv("DocServiceHost")
	// if serviceHost == "" {
	// 	log.Println("DocServiceHost не установлена в .env")
	// 	os.Exit(1)
	// }

	servicePort := os.Getenv("DocServicePort")
	if servicePort == "" {
		log.Println("DocServicePort не установлена в .env")
		os.Exit(1)
	}

	docStore := os.Getenv("DocStore")
	if docStore == "" {
		log.Println("main(): DocStore не установлена в .env")
		os.Exit(1)
	}

	sAdr := fmt.Sprintf("%v:%v", serviceHost, servicePort)
	log.Println("==================================================")
	log.Printf("[INFO] start %v %v listened: %v", serviceName, serviceID, sAdr)

	// Подключение базы данных
	err := db.NewDB()
	if err != nil {
		log.Fatalf("main(): инициализация базы данных, err=%v", err)
	}

	err = os.MkdirAll(filepath.Join(docStore), 0755)
	if err != nil {
		log.Fatalf("main(): создание папки, err=%v", err)
	}

	go func() {
		_, err := web_server.NewServer()
		if err != nil {
			log.Printf("main(): запуск сервера, err=%v", err)
			os.Exit(1)
		}
	}()

	_, err = grpc_server.NewServer(sAdr)
	if err != nil {
		log.Printf("main(): grpc_server.NewServer(): err=%v", err)
		os.Exit(1)
	}

	log.Println("main, done")
	os.Exit(0)
}
