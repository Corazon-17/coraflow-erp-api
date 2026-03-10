package middleware

import (
	"coraflow-erp-api/shared/jwt"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(jwtManager *jwt.Manager) fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Cookies("access-token")
		if token == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "access token not found")
		}

		claims, err := jwtManager.Parse(token)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
		}

		c.Locals("user-id", claims.UserID)
		c.Locals("tenant-id", claims.TenantID)

		return c.Next()
	}
}
