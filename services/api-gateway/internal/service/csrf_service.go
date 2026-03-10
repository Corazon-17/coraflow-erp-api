package service

import (
	"context"
	"time"

	"coraflow-erp-api/shared/redis"
	"coraflow-erp-api/shared/utils"
)

type CSRFService struct {
	redis      *redis.Client
	csrfTTLMin int64
}

func NewCSRFService(r *redis.Client, csrfTTLMin int64) *CSRFService {
	return &CSRFService{
		redis:      r,
		csrfTTLMin: csrfTTLMin,
	}
}

func (s *CSRFService) GenerateToken(ctx context.Context) (string, error) {
	token := utils.NewUUID().String()
	key := "csrf:" + token

	err := s.redis.Set(ctx, key, "1", time.Duration(s.csrfTTLMin)*time.Minute)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *CSRFService) ValidateToken(ctx context.Context, token string) (bool, error) {
	key := "csrf:" + token
	_, err := s.redis.Get(ctx, key)
	if err != nil {
		return false, err
	}
	return true, nil
}
