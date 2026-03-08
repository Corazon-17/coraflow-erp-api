package main

import (
	"log"

	db "coraflow-erp-api/services/user-service/db/sqlc"
	"coraflow-erp-api/services/user-service/internal/grpc"
	"coraflow-erp-api/services/user-service/internal/grpc/handler"
	"coraflow-erp-api/services/user-service/internal/repository"
	"coraflow-erp-api/services/user-service/internal/service"
	"coraflow-erp-api/shared/config"
	"coraflow-erp-api/shared/database"
)

func main() {

	cfg := config.Load()

	log.Println("starting user-service")

	pool, err := database.NewPostgres(cfg.UserDBUrl)
	if err != nil {
		log.Fatal(err)
	}

	q := db.New(pool)

	repo := repository.NewUserRepository(q)

	userService := service.NewUserService(repo)
	authService := service.NewAuthService(repo, cfg.JWTSecret)

	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	err = grpc.Start(cfg.UserServicePort, userHandler, authHandler)
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg
}
