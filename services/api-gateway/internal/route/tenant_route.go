package route

import (
	"coraflow-erp-api/services/api-gateway/internal/handler/tenant"

	"github.com/gofiber/fiber/v3"
)

func RegisterTenantRoutes(app fiber.Router, h *tenant.Handler) {

	tenantApi := app.Group("tenants")

	tenantApi.Post("", h.Create)
	tenantApi.Get("", h.List)
	tenantApi.Get("/:id", h.Get)
	tenantApi.Delete("/:id", h.Delete)

}
