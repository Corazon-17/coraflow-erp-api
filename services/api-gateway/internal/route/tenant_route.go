package route

import (
	"coraflow-erp-api/services/api-gateway/internal/handler/tenant"
	"coraflow-erp-api/services/api-gateway/internal/middleware"
	"coraflow-erp-api/shared/jwt"

	"github.com/gofiber/fiber/v3"
)

func RegisterTenantRoutes(router fiber.Router, h *tenant.Handler, jwtManager *jwt.Manager) {

	tenantApi := router.Group("tenants", middleware.AuthMiddleware(jwtManager))

	tenantApi.Post("", h.Create)
	tenantApi.Get("", h.List)
	tenantApi.Get(":id", h.Get)
	tenantApi.Delete(":id", h.Delete)

}
