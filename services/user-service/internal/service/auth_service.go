package service

import (
	"context"

	"coraflow-erp-api/services/user-service/internal/repository"
	"coraflow-erp-api/shared/jwt"
	"coraflow-erp-api/shared/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo     *repository.UserRepository
	tokenService *TokenService
}

func NewAuthService(userRepo *repository.UserRepository, tokenService *TokenService) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*jwt.Payload, error) {

	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}

	var tenantID string
	if user.TenantID != nil {
		tenantID = user.TenantID.String()
	}

	accessToken, err := s.tokenService.jwt.GenerateAccessToken(user.ID.String(), tenantID)
	if err != nil {
		return nil, nil
	}

	refreshToken, err := s.tokenService.jwt.GenerateRefreshToken(user.ID.String(), accessToken)
	if err != nil {
		return nil, nil
	}

	payload := jwt.Payload{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &payload, err
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (string, string, error) {

	userID, tokenID, err := s.tokenService.ValidateRefresh(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}

	err = s.tokenService.Revoke(ctx, tokenID)
	if err != nil {
		return "", "", err
	}

	userUUID, err := utils.ToUUID(userID)
	if err != nil {
		return "", "", err
	}

	user, err := s.userRepo.GetByID(ctx, userUUID)
	if err != nil {
		return "", "", err
	}

	var tenantID string
	if user.TenantID != nil {
		tenantID = user.TenantID.String()
	}

	return s.tokenService.GenerateTokens(ctx, user.ID.String(), tenantID)
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {

	_, tokenID, err := s.tokenService.ValidateRefresh(ctx, refreshToken)
	if err != nil {
		return err
	}

	return s.tokenService.Revoke(ctx, tokenID)
}
