package route

import (
	"coraflow-erp-api/services/api-gateway/internal/handler/user"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(router fiber.Router, auth *user.AuthHandler, userHandler *user.UserHandler) {

	authApi := router.Group("auth")
	authApi.Post("login", auth.Login)

	userApi := router.Group("users")
	userApi.Post("", userHandler.Create)

}
