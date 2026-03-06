package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"

	"coraflow-erp-api/shared/config"
)

func main() {

	cfg := config.Load()

	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "api-gateway",
			"env":     cfg.AppEnv,
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.ApiGatewayPort)))
}
