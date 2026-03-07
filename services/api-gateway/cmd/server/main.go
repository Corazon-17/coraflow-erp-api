package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"

	"coraflow-erp-api/services/api-gateway/internal/client"
	"coraflow-erp-api/services/api-gateway/internal/handler"
	"coraflow-erp-api/shared/config"
)

func main() {

	cfg := config.Load()

	app := fiber.New()

	authClient, err := client.NewAuthClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	authHandler := handler.NewAuthHandler(authClient)

	app.Post("/login", authHandler.Login)

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.ApiGatewayPort)))
}
