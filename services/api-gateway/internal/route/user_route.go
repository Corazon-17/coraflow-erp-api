package route

import (
	"coraflow-erp-api/services/api-gateway/internal/handler/user"

	"github.com/gofiber/fiber/v3"
)

func RegisterUserRoutes(app fiber.Router, auth *user.AuthHandler, userHandler *user.UserHandler) {

	authApi := app.Group("auth")
	authApi.Post("/login", auth.Login)

	userApi := app.Group("users")
	userApi.Post("/users", userHandler.Create)

}
