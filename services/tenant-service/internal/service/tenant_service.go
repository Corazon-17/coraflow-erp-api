package service

import (
	"context"

	db "coraflow-erp-api/services/tenant-service/db/sqlc"
	"coraflow-erp-api/services/tenant-service/internal/repository"
	"coraflow-erp-api/shared/utils"

	"github.com/google/uuid"
)

type TenantService struct {
	repo *repository.TenantRepository
}

func NewTenantService(r *repository.TenantRepository) *TenantService {
	return &TenantService{repo: r}
}

func (s *TenantService) CreateTenant(ctx context.Context, name string, slug string) (*db.Tenant, error) {

	t, err := s.repo.Create(ctx, db.CreateTenantParams{
		ID:   utils.NewUUID(),
		Name: name,
		Slug: slug,
	})

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *TenantService) GetTenant(ctx context.Context, id string) (*db.Tenant, error) {

	tenantUUID, err := utils.ToUUID(id)
	if err != nil {
		return nil, err
	}

	t, err := s.repo.GetByID(ctx, tenantUUID)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *TenantService) GetTenantBySlug(ctx context.Context, slug string) (*db.Tenant, error) {

	t, err := s.repo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *TenantService) ListTenants(ctx context.Context) ([]db.Tenant, error) {

	return s.repo.List(ctx)
}

func (s *TenantService) UpdateTenant(ctx context.Context, id string, name string, slug string) (*db.Tenant, error) {

	tenantUUID, err := utils.ToUUID(id)
	if err != nil {
		return nil, err
	}

	t, err := s.repo.Update(ctx, db.UpdateTenantParams{
		ID:   tenantUUID,
		Name: name,
		Slug: slug,
	})

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *TenantService) DeleteTenant(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
