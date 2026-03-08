package handler

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

func (h *DepartmentHandler) CreateDepartment(c fiber.Ctx) error {

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
