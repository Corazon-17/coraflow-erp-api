package service

import (
	"context"

	"coraflow-erp-api/services/hr-service/internal/repository"

	"github.com/google/uuid"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) GetEmployee(ctx context.Context, tenantID, employeeID uuid.UUID) (interface{}, error) {

	return s.repo.GetEmployee(ctx, tenantID, employeeID)
}