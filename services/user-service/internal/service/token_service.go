package service

import (
	"context"
	"fmt"
	"time"

	"coraflow-erp-api/shared/jwt"
	"coraflow-erp-api/shared/redis"
	"coraflow-erp-api/shared/utils"
)

type TokenService struct {
	jwt              *jwt.Manager
	redis            *redis.Client
	jwtRefreshTTLMin int64
}

func NewTokenService(jwt *jwt.Manager, redis *redis.Client, jwtRefreshTTLMin int64) *TokenService {
	return &TokenService{
		jwt:              jwt,
		redis:            redis,
		jwtRefreshTTLMin: jwtRefreshTTLMin,
	}
}

func (s *TokenService) GenerateTokens(ctx context.Context, userID string, tenantID string) (string, string, error) {

	tokenID := utils.NewUUID().String()

	access, err := s.jwt.GenerateAccessToken(userID, tenantID)
	if err != nil {
		return "", "", err
	}

	refresh, err := s.jwt.GenerateRefreshToken(userID, tokenID)
	if err != nil {
		return "", "", err
	}

	key := fmt.Sprintf("refresh:%s", tokenID)

	err = s.redis.Set(ctx, key, userID, time.Duration(s.jwtRefreshTTLMin)*time.Minute)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *TokenService) ValidateRefresh(ctx context.Context, token string) (string, string, error) {

	claims, err := s.jwt.Parse(token)
	if err != nil {
		return "", "", err
	}

	key := fmt.Sprintf("refresh:%s", claims.TokenID)

	userID, err := s.redis.Get(ctx, key)
	if err != nil {
		return "", "", err
	}

	return userID, claims.TokenID, nil
}

func (s *TokenService) Revoke(ctx context.Context, tokenID string) error {

	key := fmt.Sprintf("refresh:%s", tokenID)

	return s.redis.Delete(ctx, key)
}
