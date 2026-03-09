package user

import (
	"context"

	userpb "coraflow-erp-api/proto/user/user/v1"
	"coraflow-erp-api/services/api-gateway/internal/client"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	client *client.UserClient
}

func NewUserHandler(c *client.UserClient) *UserHandler {
	return &UserHandler{client: c}
}

func (h *UserHandler) Create(c fiber.Ctx) error {

	req := new(userpb.CreateUserRequest)

	if err := c.Bind().Body(req); err != nil {
		return err
	}

	res, err := h.client.User.CreateUser(context.Background(), req)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
