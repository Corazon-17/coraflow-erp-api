package route

import (
	"coraflow-erp-api/services/api-gateway/internal/handler"

	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(
	app *fiber.App,
	auth *handler.AuthHandler,
	tenant *handler.TenantHandler,
	department *handler.DepartmentHandler,
	employee *handler.EmployeeHandler,
) {

	app.Post("/login", auth.Login)
	app.Post("/tenants", tenant.CreateTenant)
	app.Post("/departments", department.CreateDepartment)
	app.Post("/employees", employee.CreateEmployee)
}
