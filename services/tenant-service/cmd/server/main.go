package main

import (
	"log"

	db "coraflow-erp-api/services/tenant-service/db/sqlc"
	"coraflow-erp-api/services/tenant-service/internal/grpc"
	"coraflow-erp-api/services/tenant-service/internal/grpc/handler"
	"coraflow-erp-api/services/tenant-service/internal/repository"
	"coraflow-erp-api/services/tenant-service/internal/service"
	"coraflow-erp-api/shared/config"
	"coraflow-erp-api/shared/database"
)

func main() {

	cfg := config.Load()

	log.Println("starting tenant-service")

	pool, err := database.NewPostgres(cfg.TenantDBUrl)
	if err != nil {
		log.Fatal(err)
	}

	q := db.New(pool)

	repo := repository.NewTenantRepository(q)

	service := service.NewTenantService(repo)

	handler := handler.NewTenantHandler(service)

	err = grpc.Start(cfg.TenantServicePort, handler)
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg
}
