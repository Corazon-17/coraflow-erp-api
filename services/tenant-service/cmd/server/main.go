package main

import (
	"fmt"
	"log"

	"coraflow-erp-api/services/tenant-service/internal/grpc"
	"coraflow-erp-api/shared/config"
)

func main() {

	cfg := config.Load()

	log.Println("starting tenant-service")

	err := grpc.Start(fmt.Sprintf(":%s", cfg.TenantServicePort))
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg
}
