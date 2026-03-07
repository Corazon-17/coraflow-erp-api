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

func (r *EmployeeRepository) GetEmployee(ctx context.Context, tenantID, id uuid.UUID) (db.Employee, error) {

	return r.q.GetEmployee(ctx, db.GetEmployeeParams{
		TenantID: tenantID,
		ID: id,
	})
}