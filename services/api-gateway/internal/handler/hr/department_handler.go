package hr

import (
	"context"

	departmentpb "coraflow-erp-api/proto/hr/department/v1"
	"coraflow-erp-api/services/api-gateway/internal/client"

	"github.com/gofiber/fiber/v3"
)

type DepartmentHandler struct {
	client *client.HRClient
}

func NewDepartmentHandler(c *client.HRClient) *DepartmentHandler {
	return &DepartmentHandler{client: c}
}

func (h *DepartmentHandler) Create(c fiber.Ctx) error {

	req := new(departmentpb.CreateDepartmentRequest)

	if err := c.Bind().Body(req); err != nil {
		return err
	}

	res, err := h.client.Department.CreateDepartment(context.Background(), req)
	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *DepartmentHandler) Get(c fiber.Ctx) error {

	id := c.Params("id")
	tenant := c.Query("tenant_id")

	res, err := h.client.Department.GetDepartment(context.Background(),
		&departmentpb.GetDepartmentRequest{
			Id:       id,
			TenantId: tenant,
		})

	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *DepartmentHandler) List(c fiber.Ctx) error {

	tenant := c.Query("tenant_id")

	res, err := h.client.Department.ListDepartment(context.Background(),
		&departmentpb.ListDepartmentRequest{
			TenantId: tenant,
		})

	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *DepartmentHandler) Delete(c fiber.Ctx) error {

	id := c.Params("id")
	tenant := c.Query("tenant_id")

	res, err := h.client.Department.DeleteDepartment(context.Background(),
		&departmentpb.DeleteDepartmentRequest{
			Id:       id,
			TenantId: tenant,
		})

	if err != nil {
		return err
	}

	return c.JSON(res)
}
