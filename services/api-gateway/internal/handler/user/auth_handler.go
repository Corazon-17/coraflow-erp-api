package user

import (
	"context"

	authpb "coraflow-erp-api/proto/user/auth/v1"
	"coraflow-erp-api/services/api-gateway/internal/client"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	client *client.UserClient
}

func NewAuthHandler(c *client.UserClient) *AuthHandler {
	return &AuthHandler{client: c}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {

	req := new(authpb.LoginRequest)
	if err := c.Bind().Body(req); err != nil {
		return err
	}

	res, err := h.client.Auth.Login(context.Background(), req)
	if err != nil {
		return err
	}

	access := res.AccessToken
	refresh := res.RefreshToken

	c.Cookie(&fiber.Cookie{
		Name:     "access-token",
		Value:    access,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refresh-token",
		Value:    refresh,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/auth/refresh",
	})

	return c.JSON(fiber.Map{
		"message": "login success",
	})
}
