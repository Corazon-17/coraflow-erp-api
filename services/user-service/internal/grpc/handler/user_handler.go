package handler

import (
	"context"

	userpb "coraflow-erp-api/proto/user/user/v1"
	"coraflow-erp-api/services/user-service/internal/service"
	"coraflow-erp-api/shared/utils"

	"github.com/google/uuid"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.UserResponse, error) {

	var tenantUUID *uuid.UUID
	if req.TenantId != "" {
		id, err := utils.ToUUID(req.TenantId)
		if err != nil {
			return nil, err
		}
		tenantUUID = &id
	}

	u, err := h.service.CreateUser(ctx, req.Email, req.Password, tenantUUID, req.IsInternal)
	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{
		Id:    u.ID.String(),
		Email: u.Email,
	}, nil
}
