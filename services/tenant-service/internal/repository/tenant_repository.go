package repository

import (
	"context"

	db "coraflow-erp-api/services/tenant-service/db/sqlc"
)

type TenantRepository struct {
	q *db.Queries
}

func NewTenantRepository(q *db.Queries) *TenantRepository {
	return &TenantRepository{q: q}
}

func (r *TenantRepository) CreateTenant(ctx context.Context, arg db.CreateTenantParams) (db.Tenant, error) {
	return r.q.CreateTenant(ctx, arg)
}