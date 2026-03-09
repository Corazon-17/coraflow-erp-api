package main

import (
	"log"

	"coraflow-erp-api/services/api-gateway/internal/client"
	"coraflow-erp-api/services/api-gateway/internal/handler/hr"
	"coraflow-erp-api/services/api-gateway/internal/handler/tenant"
	"coraflow-erp-api/services/api-gateway/internal/handler/user"
	"coraflow-erp-api/services/api-gateway/internal/route"
	"coraflow-erp-api/shared/config"
	"coraflow-erp-api/shared/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {

	cfg := config.Load()

	app := fiber.New(fiber.Config{
		StrictRouting: false,
	})
	app.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowOrigins: []string{"*"},
	}))

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

	tenantHandler := tenant.NewTenantHandler(tenantClient)
	authHandler := user.NewAuthHandler(userClient)
	userHandler := user.NewUserHandler(userClient)
	deptHandler := hr.NewDepartmentHandler(hrClient)
	empHandler := hr.NewEmployeeHandler(hrClient)

	api := app.Group("api")

	route.RegisterTenantRoutes(api, tenantHandler)
	route.RegisterUserRoutes(api, authHandler, userHandler)
	route.RegisterHRRoutes(api, deptHandler, empHandler)

	log.Fatal(app.Listen(utils.GetPort(cfg.ApiGatewayPort)))
}
