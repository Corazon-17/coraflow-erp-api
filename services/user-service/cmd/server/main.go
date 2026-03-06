package main

import (
	"fmt"
	"log"

	"coraflow-erp-api/services/user-service/internal/grpc"
	"coraflow-erp-api/shared/config"
)

func main() {

	cfg := config.Load()

	log.Println("starting user-service")

	err := grpc.Start(fmt.Sprintf(":%s", cfg.UserServicePort))
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg
}
