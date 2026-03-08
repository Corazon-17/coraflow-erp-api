package handler

import (
	"context"

	tenantpb "coraflow-erp-api/proto/tenant/tenant/v1"
	"coraflow-erp-api/services/api-gateway/internal/client"

	"github.com/gofiber/fiber/v3"
)

type TenantHandler struct {
	client *client.TenantClient
}

func NewTenantHandler(c *client.TenantClient) *TenantHandler {
	return &TenantHandler{client: c}
}

func (h *TenantHandler) CreateTenant(c fiber.Ctx) error {

	req := new(tenantpb.CreateTenantRequest)

	if err := c.Bind().Body(req); err != nil {
		return err
	}

	res, err := h.client.Client.CreateTenant(context.Background(), req)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
