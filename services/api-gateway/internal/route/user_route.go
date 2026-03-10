package route

import (
	"coraflow-erp-api/services/api-gateway/internal/handler/user"
	"coraflow-erp-api/services/api-gateway/internal/middleware"
	"coraflow-erp-api/shared/jwt"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(router fiber.Router, authHandler *user.AuthHandler, userHandler *user.UserHandler, jwtManager *jwt.Manager) {

	authApi := router.Group("auth")
	authApi.Post("login", authHandler.Login)

	userApi := router.Group("users", middleware.AuthMiddleware(jwtManager))
	userApi.Post("", userHandler.Create)

}
