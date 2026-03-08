package service

import (
	"context"

	db "coraflow-erp-api/services/hr-service/db/sqlc"
	"coraflow-erp-api/services/hr-service/internal/repository"
	"coraflow-erp-api/shared/utils"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) Create(ctx context.Context, tenantID string, first string, last string) (*db.Employee, error) {

	tenantUUID, err := utils.ToUUID(tenantID)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.Create(ctx, db.CreateEmployeeParams{
		ID:        utils.NewUUID(),
		TenantID:  tenantUUID,
		FirstName: first,
		LastName:  last,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *EmployeeService) Get(ctx context.Context, tenantID string, id string) (*db.Employee, error) {

	tenantUUID, err := utils.ToUUID(tenantID)
	if err != nil {
		return nil, err
	}

	empUUID, err := utils.ToUUID(id)
	if err != nil {
		return nil, err
	}

	result, err := s.repo.Get(ctx, db.GetEmployeeParams{
		TenantID: tenantUUID,
		ID:       empUUID,
	})
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *EmployeeService) List(ctx context.Context, tenantID string) ([]db.Employee, error) {

	tenantUUID, err := utils.ToUUID(tenantID)
	if err != nil {
		return nil, err
	}

	return s.repo.List(ctx, tenantUUID)
}

func (s *EmployeeService) Delete(ctx context.Context, tenantID string, id string) error {

	tenantUUID, err := utils.ToUUID(tenantID)
	if err != nil {
		return err
	}

	empUUID, err := utils.ToUUID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, tenantUUID, empUUID)
}
