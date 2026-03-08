package service

import (
	"context"

	db "coraflow-erp-api/services/hr-service/db/sqlc"
	"coraflow-erp-api/services/hr-service/internal/repository"
	"coraflow-erp-api/shared/utils"

	"github.com/google/uuid"
)

type DepartmentService struct {
	repo *repository.DepartmentRepository
}

func NewDepartmentService(repo *repository.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (s *DepartmentService) Create(ctx context.Context, tenantID string, name string, parentID *string) (*db.Department, error) {

	tenantUUID, err := utils.ToUUID(tenantID)
	if err != nil {
		return nil, err
	}

	var parentUUID *uuid.UUID

	if parentID != nil && *parentID != "" {
		id, err := utils.ToUUID(*parentID)
		if err != nil {
			return nil, err
		}
		parentUUID = &id
	}

	result, err := s.repo.Create(ctx, db.CreateDepartmentParams{
		ID:       utils.NewUUID(),
		TenantID: tenantUUID,
		Name:     name,
		ParentID: parentUUID,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DepartmentService) Get(ctx context.Context, tenantID string, id string) (*db.Department, error) {

	tenantUUID, err := utils.ToUUID(tenantID)
	if err != nil {
		return nil, err
	}

	deptUUID, err := utils.ToUUID(id)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.Get(ctx, tenantUUID, deptUUID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *DepartmentService) List(ctx context.Context, tenantID string) ([]db.Department, error) {

	tenantUUID, err := utils.ToUUID(tenantID)
	if err != nil {
		return nil, err
	}

	return s.repo.List(ctx, tenantUUID)
}

func (s *DepartmentService) Delete(ctx context.Context, tenantID string, id string) error {

	tenantUUID, err := utils.ToUUID(tenantID)
	if err != nil {
		return err
	}

	deptUUID, err := utils.ToUUID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, tenantUUID, deptUUID)
}
