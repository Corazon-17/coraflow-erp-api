package hr

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

func (h *EmployeeHandler) Create(c fiber.Ctx) error {

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

func (h *EmployeeHandler) Get(c fiber.Ctx) error {

	id := c.Params("id")
	tenant := c.Query("tenant_id")

	res, err := h.client.Employee.GetEmployee(context.Background(),
		&employeepb.GetEmployeeRequest{
			Id:       id,
			TenantId: tenant,
		})

	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *EmployeeHandler) List(c fiber.Ctx) error {

	tenant := c.Query("tenant_id")

	res, err := h.client.Employee.ListEmployee(context.Background(),
		&employeepb.ListEmployeeRequest{
			TenantId: tenant,
		})

	if err != nil {
		return err
	}

	return c.JSON(res)
}

func (h *EmployeeHandler) Delete(c fiber.Ctx) error {

	id := c.Params("id")
	tenant := c.Query("tenant_id")

	res, err := h.client.Employee.DeleteEmployee(context.Background(),
		&employeepb.DeleteEmployeeRequest{
			Id:       id,
			TenantId: tenant,
		})

	if err != nil {
		return err
	}

	return c.JSON(res)
}
