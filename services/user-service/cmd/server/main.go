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
	"coraflow-erp-api/shared/jwt"
	"coraflow-erp-api/shared/redis"
)

func main() {

	cfg := config.Load()

	log.Println("starting user-service")

	pool, err := database.NewPostgres(cfg.UserDBUrl)
	if err != nil {
		log.Fatal(err)
	}

	q := db.New(pool)

	jwt := jwt.New(cfg.JWTSecret)

	rds := redis.NewRedis(cfg.RabbitMQUrl)

	repo := repository.NewUserRepository(q)

	jwtService := service.NewTokenService(jwt, rds)
	userService := service.NewUserService(repo)
	authService := service.NewAuthService(repo, jwtService)

	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	err = grpc.Start(cfg.UserServicePort, userHandler, authHandler)
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg
}
