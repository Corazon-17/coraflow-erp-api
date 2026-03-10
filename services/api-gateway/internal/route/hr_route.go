package route

import (
	"coraflow-erp-api/services/api-gateway/internal/handler/hr"
	"coraflow-erp-api/services/api-gateway/internal/middleware"
	"coraflow-erp-api/shared/jwt"

	"github.com/gofiber/fiber/v3"
)

func RegisterHRRoutes(router fiber.Router, dept *hr.DepartmentHandler, emp *hr.EmployeeHandler, jwtManager *jwt.Manager) {

	department := router.Group("departments", middleware.AuthMiddleware(jwtManager))

	department.Post("", dept.Create)
	department.Get("", dept.List)
	department.Get(":id", dept.Get)
	department.Delete(":id", dept.Delete)

	employee := router.Group("employees", middleware.AuthMiddleware(jwtManager))

	employee.Post("", emp.Create)
	employee.Get("", emp.List)
	employee.Get(":id", emp.Get)
	employee.Delete(":id", emp.Delete)
}
