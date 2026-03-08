package repository

import (
	"context"

	db "coraflow-erp-api/services/tenant-service/db/sqlc"

	"github.com/google/uuid"
)

type TenantRepository struct {
	q *db.Queries
}

func NewTenantRepository(q *db.Queries) *TenantRepository {
	return &TenantRepository{q: q}
}

func (r *TenantRepository) Create(ctx context.Context, arg db.CreateTenantParams) (db.Tenant, error) {
	return r.q.CreateTenant(ctx, arg)
}

func (r *TenantRepository) GetByID(ctx context.Context, id uuid.UUID) (db.Tenant, error) {
	return r.q.GetTenantByID(ctx, id)
}

func (r *TenantRepository) GetBySlug(ctx context.Context, slug string) (db.Tenant, error) {
	return r.q.GetTenantBySlug(ctx, slug)
}

func (r *TenantRepository) List(ctx context.Context) ([]db.Tenant, error) {
	return r.q.ListTenant(ctx)
}

func (r *TenantRepository) Update(ctx context.Context, arg db.UpdateTenantParams) (db.Tenant, error) {
	return r.q.UpdateTenant(ctx, arg)
}

func (r *TenantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTenant(ctx, id)
}
