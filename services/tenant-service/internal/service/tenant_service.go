package service

import (
	"context"

	db "coraflow-erp-api/services/tenant-service/db/sqlc"
	"coraflow-erp-api/services/tenant-service/internal/repository"
	"coraflow-erp-api/shared/utils"
)

type TenantService struct {
	repo *repository.TenantRepository
}

func NewTenantService(repo *repository.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (s *TenantService) CreateTenant(ctx context.Context, name string, slug string) (db.Tenant, error) {

	return s.repo.CreateTenant(ctx, db.CreateTenantParams{
		ID:   utils.NewID(),
		Name: name,
		Slug: slug,
	})
}