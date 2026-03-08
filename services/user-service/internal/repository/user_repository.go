package repository

import (
	"context"

	db "coraflow-erp-api/services/user-service/db/sqlc"

	"github.com/google/uuid"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) Create(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.q.CreateUser(ctx, arg)
}

func (r *UserRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (db.User, error) {
	return r.q.GetUserByID(ctx, id)
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (db.User, error) {
	return r.q.GetUserByEmail(ctx, email)
}

func (r *UserRepository) ListByTenant(ctx context.Context, tenantID *uuid.UUID) ([]db.User, error) {
	return r.q.ListUserByTenant(ctx, tenantID)
}

func (r *UserRepository) ListInternal(ctx context.Context) ([]db.User, error) {
	return r.q.ListInternalUser(ctx)
}

func (r *UserRepository) UpdateTenant(ctx context.Context, id uuid.UUID, tenantID *uuid.UUID) (db.User, error) {
	return r.q.UpdateUserTenant(ctx, db.UpdateUserTenantParams{
		ID:       id,
		TenantID: tenantID,
	})
}

func (r *UserRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return r.q.DeleteUser(ctx, id)
}
