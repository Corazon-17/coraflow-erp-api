package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"

	authpb "coraflow-erp-api/proto/user/auth/v1"
	"coraflow-erp-api/services/api-gateway/internal/client"
	"coraflow-erp-api/shared/config"
)

func main() {

	cfg := config.Load()

	app := fiber.New()

	authClient, err := client.NewAuthClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "api-gateway",
			"env":     cfg.AppEnv,
		})
	})

	app.Post("/login", func(c fiber.Ctx) error {

		req := new(struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		})

		if err := c.Bind().Body(req); err != nil {
			return err
		}

		res, err := authClient.Login(context.Background(), &authpb.LoginRequest{
			Email:    req.Email,
			Password: req.Password,
		})

		if err != nil {
			return err
		}

		return c.JSON(res)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.ApiGatewayPort)))
}
