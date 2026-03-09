package route

import (
	"coraflow-erp-api/services/api-gateway/internal/handler/hr"

	"github.com/gofiber/fiber/v3"
)

func RegisterHRRoutes(
	app fiber.Router,
	dept *hr.DepartmentHandler,
	emp *hr.EmployeeHandler,
) {

	department := app.Group("/departments")

	department.Post("/", dept.Create)
	department.Get("/", dept.List)
	department.Get("/:id", dept.Get)
	department.Delete("/:id", dept.Delete)

	employee := app.Group("/employees")

	employee.Post("/", emp.Create)
	employee.Get("/", emp.List)
	employee.Get("/:id", emp.Get)
	employee.Delete("/:id", emp.Delete)
}
