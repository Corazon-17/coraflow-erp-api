package main

import (
	"fmt"
	"log"

	"coraflow-erp-api/services/hr-service/internal/grpc"
	"coraflow-erp-api/shared/config"
)

func main() {
	cfg := config.Load()

	log.Println("starting hr-service")

	err := grpc.Start(fmt.Sprintf(":%s", cfg.HRServicePort))
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg

}
