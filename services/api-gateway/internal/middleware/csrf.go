package middleware

import (
	"context"
	"coraflow-erp-api/shared/redis"

	"github.com/gofiber/fiber/v3"
)

func CSRFMiddleware(r *redis.Client) fiber.Handler {
	return func(c fiber.Ctx) error {

		method := c.Method()
		if method != fiber.MethodPost &&
			method != fiber.MethodPut &&
			method != fiber.MethodPatch &&
			method != fiber.MethodDelete {
			return c.Next()
		}

		csrfCookie := c.Cookies("csrf-token")
		csrfHeader := c.Get("X-CSRF-Token")

		if csrfCookie == "" || csrfHeader == "" {
			return fiber.NewError(fiber.StatusForbidden, "CSRF token missing")
		}

		if csrfCookie != csrfHeader {
			return fiber.NewError(fiber.StatusForbidden, "CSRF token mismatch")
		}

		ok, err := r.Get(context.Background(), "csrf:"+csrfHeader)
		if err != nil || ok == "" {
			return fiber.NewError(fiber.StatusForbidden, "CSRF token invalid")
		}

		return c.Next()
	}
}
