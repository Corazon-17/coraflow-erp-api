package service

import (
	"context"

	db "coraflow-erp-api/services/user-service/db/sqlc"
	"coraflow-erp-api/services/user-service/internal/repository"
	"coraflow-erp-api/shared/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(ctx context.Context, email string, password string, tenantID *uuid.UUID, isInternal bool) (*db.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Create(ctx, db.CreateUserParams{
		ID:           utils.NewUUID(),
		Email:        email,
		PasswordHash: string(hash),
		TenantID:     tenantID,
		IsInternal:   isInternal,
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}
