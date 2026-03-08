package repository

import (
	"context"

	db "coraflow-erp-api/services/hr-service/db/sqlc"

	"github.com/google/uuid"
)

type EmployeeRepository struct {
	q *db.Queries
}

func NewEmployeeRepository(q *db.Queries) *EmployeeRepository {
	return &EmployeeRepository{q: q}
}

func (r *EmployeeRepository) Create(ctx context.Context, arg db.CreateEmployeeParams) (db.Employee, error) {
	return r.q.CreateEmployee(ctx, arg)
}

func (r *EmployeeRepository) Get(ctx context.Context, arg db.GetEmployeeParams) (db.Employee, error) {
	return r.q.GetEmployee(ctx, arg)
}

func (r *EmployeeRepository) GetByUser(ctx context.Context, arg db.GetEmployeeByUserParams) (db.Employee, error) {
	return r.q.GetEmployeeByUser(ctx, arg)
}

func (r *EmployeeRepository) List(ctx context.Context, tenantID uuid.UUID) ([]db.Employee, error) {
	return r.q.ListEmployee(ctx, tenantID)
}

func (r *EmployeeRepository) Update(ctx context.Context, arg db.UpdateEmployeeParams) (db.Employee, error) {
	return r.q.UpdateEmployee(ctx, arg)
}

func (r *EmployeeRepository) Delete(ctx context.Context, tenantID uuid.UUID, id uuid.UUID) error {
	return r.q.DeleteEmployee(ctx, db.DeleteEmployeeParams{
		TenantID: tenantID,
		ID:       id,
	})
}
