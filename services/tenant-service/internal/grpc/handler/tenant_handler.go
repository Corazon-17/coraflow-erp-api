package handlers

import (
	"context"

	tenantpb "coraflow-erp-api/proto/tenant/tenant/v1"
	"coraflow-erp-api/services/tenant-service/internal/service"
)

type TenantHandler struct {
	tenantpb.UnimplementedTenantServiceServer
	service *service.TenantService
}

func NewTenantHandler(s *service.TenantService) *TenantHandler {
	return &TenantHandler{service: s}
}

func (h *TenantHandler) CreateTenant(ctx context.Context, req *tenantpb.CreateTenantRequest) (*tenantpb.TenantResponse, error) {

	t, err := h.service.CreateTenant(ctx, req.Name, req.Slug)
	if err != nil {
		return nil, err
	}

	return &tenantpb.TenantResponse{
	Tenant: &tenantpb.Tenant{
		Id:   t.ID.String(),
		Name: t.Name,
		Slug: t.Slug,
	},
}, nil
}