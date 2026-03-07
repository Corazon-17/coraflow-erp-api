package handler

import (
	"context"

	authpb "coraflow-erp-api/proto/user/auth/v1"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	client authpb.AuthServiceClient
}

func NewAuthHandler(c authpb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{client: c}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {

	req := new(struct {
		Email string `json:"email"`
		Password string `json:"password"`
	})

	if err := c.Bind().Body(req); err != nil {
		return err
	}

	res, err := h.client.Login(context.Background(), &authpb.LoginRequest{
		Email: req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	return c.JSON(res)
}