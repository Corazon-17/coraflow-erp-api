package main

import (
	"log"

	"coraflow-erp-api/services/api-gateway/internal/client"
	"coraflow-erp-api/services/api-gateway/internal/handler/hr"
	"coraflow-erp-api/services/api-gateway/internal/handler/tenant"
	"coraflow-erp-api/services/api-gateway/internal/handler/user"
	"coraflow-erp-api/services/api-gateway/internal/middleware"
	"coraflow-erp-api/services/api-gateway/internal/route"
	"coraflow-erp-api/services/api-gateway/internal/service"
	"coraflow-erp-api/shared/config"
	"coraflow-erp-api/shared/jwt"
	"coraflow-erp-api/shared/redis"
	"coraflow-erp-api/shared/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {

	cfg := config.Load()
	rds := redis.NewRedis(cfg.RedisUrl)

	app := fiber.New(fiber.Config{
		StrictRouting: false,
	})
	app.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowOrigins: []string{"http://localhost:3000"},
	}))
	app.Use(middleware.CSRFMiddleware(rds))

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

	csrfService := service.NewCSRFService(rds, cfg.CSRFTTLMin)

	tenantHandler := tenant.NewTenantHandler(tenantClient)
	authHandler := user.NewAuthHandler(userClient, csrfService)
	userHandler := user.NewUserHandler(userClient)
	deptHandler := hr.NewDepartmentHandler(hrClient)
	empHandler := hr.NewEmployeeHandler(hrClient)

	jwtManager := jwt.New(cfg.JWTSecret, cfg.JWTAccessTTLMin, cfg.JWTRefreshTTLMin)

	api := app.Group("api")

	route.RegisterTenantRoutes(api, tenantHandler, jwtManager)
	route.RegisterUserRoutes(api, authHandler, userHandler, jwtManager)
	route.RegisterHRRoutes(api, deptHandler, empHandler, jwtManager)

	log.Fatal(app.Listen(utils.GetPort(cfg.ApiGatewayPort)))
}
