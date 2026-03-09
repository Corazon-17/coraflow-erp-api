package tenant

import (
	"context"

	tenantpb "coraflow-erp-api/proto/tenant/tenant/v1"
	"coraflow-erp-api/services/api-gateway/internal/client"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	client *client.TenantClient
}

func NewTenantHandler(c *client.TenantClient) *Handler {
	return &Handler{client: c}
}

func (h *Handler) Create(c fiber.Ctx) error {

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

func (h *Handler) Get(c fiber.Ctx) error {

	id := c.Params("id")

	res, err := h.client.Client.GetTenant(context.Background(),
		&tenantpb.GetTenantRequest{
			Id: id,
		})
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *Handler) List(c fiber.Ctx) error {

	res, err := h.client.Client.ListTenants(context.Background(), &tenantpb.ListTenantRequest{})
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *Handler) Delete(c fiber.Ctx) error {

	id := c.Params("id")

	res, err := h.client.Client.DeleteTenant(context.Background(),
		&tenantpb.DeleteTenantRequest{
			Id: id,
		})
	if err != nil {
		return err
	}

	return c.JSON(res)
}
