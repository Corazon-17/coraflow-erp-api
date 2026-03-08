package handler

import (
	"context"

	employeepb "coraflow-erp-api/proto/hr/employee/v1"
	"coraflow-erp-api/services/api-gateway/internal/client"

	"github.com/gofiber/fiber/v3"
)

type EmployeeHandler struct {
	client *client.HRClient
}

func NewEmployeeHandler(c *client.HRClient) *EmployeeHandler {
	return &EmployeeHandler{client: c}
}

func (h *EmployeeHandler) CreateEmployee(c fiber.Ctx) error {

	req := new(employeepb.CreateEmployeeRequest)

	if err := c.Bind().Body(req); err != nil {
		return err
	}

	res, err := h.client.Employee.CreateEmployee(context.Background(), req)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
