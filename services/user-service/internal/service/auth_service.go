package service

import (
	"context"

	"coraflow-erp-api/services/user-service/internal/repository"
	"coraflow-erp-api/shared/jwt"
	"coraflow-erp-api/shared/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
	jwtSecret string
}

func NewAuthService(repo *repository.UserRepository, secret string) *AuthService {
	return &AuthService{
		repo: repo,
		jwtSecret: secret,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", err
	}

	var tenantID *string
	if utils.IsValidUUID(user.TenantID.String()) {
		id := user.TenantID.String()
		tenantID = &id
	}

	token, err := jwt.Generate(s.jwtSecret, jwt.Claims{
		UserID: user.ID.String(),
		TenantID: tenantID,
		IsInternal: user.IsInternal,
	})

	return token, err
}