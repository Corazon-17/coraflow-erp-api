package main

import (
	"fmt"
	"log"

	"coraflow-erp-api/services/api-gateway/internal/client"
	"coraflow-erp-api/services/api-gateway/internal/handler"
	"coraflow-erp-api/services/api-gateway/internal/route"
	"coraflow-erp-api/shared/config"

	"github.com/gofiber/fiber/v3"
)

func main() {

	cfg := config.Load()

	app := fiber.New()

	tenantClient, err := client.NewTenantClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	userClient, err := client.NewUserClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	hrClient, err := client.NewHRClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	authHandler := handler.NewAuthHandler(userClient)
	tenantHandler := handler.NewTenantHandler(tenantClient)
	departmentHandler := handler.NewDepartmentHandler(hrClient)
	employeeHandler := handler.NewEmployeeHandler(hrClient)

	route.RegisterRoutes(
		app,
		authHandler,
		tenantHandler,
		departmentHandler,
		employeeHandler,
	)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", cfg.ApiGatewayPort)))
}
