package user

import (
	"context"

	authpb "coraflow-erp-api/proto/user/auth/v1"
	"coraflow-erp-api/services/api-gateway/internal/client"
	"coraflow-erp-api/services/api-gateway/internal/service"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	client *client.UserClient
	csrf   *service.CSRFService
}

func NewAuthHandler(c *client.UserClient, csrf *service.CSRFService) *AuthHandler {
	return &AuthHandler{
		client: c,
		csrf:   csrf,
	}
}

func (h *AuthHandler) CSRF(c fiber.Ctx) error {

	token, err := h.csrf.GenerateToken(context.Background())
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "csrf_token",
		Value:    token,
		HTTPOnly: false,
		Secure:   false,
		SameSite: "Strict",
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"csrfToken": token,
	})
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
