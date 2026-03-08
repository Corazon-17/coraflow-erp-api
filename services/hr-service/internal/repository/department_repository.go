package repository

import (
	"context"

	db "coraflow-erp-api/services/hr-service/db/sqlc"

	"github.com/google/uuid"
)

type DepartmentRepository struct {
	q *db.Queries
}

func NewDepartmentRepository(q *db.Queries) *DepartmentRepository {
	return &DepartmentRepository{q: q}
}

func (r *DepartmentRepository) Create(ctx context.Context, arg db.CreateDepartmentParams) (db.Department, error) {

	return r.q.CreateDepartment(ctx, arg)
}

func (r *DepartmentRepository) Get(ctx context.Context, tenantID uuid.UUID, id uuid.UUID) (db.Department, error) {

	return r.q.GetDepartment(ctx, db.GetDepartmentParams{
		TenantID: tenantID,
		ID:       id,
	})
}

func (r *DepartmentRepository) List(ctx context.Context, tenantID uuid.UUID) ([]db.Department, error) {
	return r.q.ListDepartment(ctx, tenantID)
}

func (r *DepartmentRepository) Update(ctx context.Context, arg db.UpdateDepartmentParams) (db.Department, error) {
	return r.q.UpdateDepartment(ctx, arg)
}

func (r *DepartmentRepository) Delete(ctx context.Context, tenantID uuid.UUID, id uuid.UUID) error {

	return r.q.DeleteDepartment(ctx, db.DeleteDepartmentParams{
		TenantID: tenantID,
		ID:       id,
	})
}
