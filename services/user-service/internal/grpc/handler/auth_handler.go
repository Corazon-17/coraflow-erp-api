package handler

import (
	"context"

	authpb "coraflow-erp-api/proto/user/auth/v1"
	"coraflow-erp-api/services/user-service/internal/service"
)

type AuthHandler struct {
	authpb.UnimplementedAuthServiceServer
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	token, err := h.service.Login(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
