package repository

import (
	"context"

	db "coraflow-erp-api/services/user-service/db/sqlc"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return r.q.GetUserByEmail(ctx, email)
}