package handler

import (
	"context"

	employeepb "coraflow-erp-api/proto/hr/employee/v1"
	"coraflow-erp-api/services/hr-service/internal/service"
)

type EmployeeHandler struct {
	employeepb.UnimplementedEmployeeServiceServer
	service *service.EmployeeService
}

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		service: s,
	}
}

func (h *EmployeeHandler) CreateEmployee(ctx context.Context, req *employeepb.CreateEmployeeRequest) (*employeepb.EmployeeResponse, error) {

	e, err := h.service.Create(
		ctx,
		req.TenantId,
		req.FirstName,
		req.LastName,
	)
	if err != nil {
		return nil, err
	}

	return &employeepb.EmployeeResponse{
		Id:        e.ID.String(),
		TenantId:  e.TenantID.String(),
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}, nil
}

func (h *EmployeeHandler) GetEmployee(ctx context.Context, req *employeepb.GetEmployeeRequest) (*employeepb.EmployeeResponse, error) {

	e, err := h.service.Get(
		ctx,
		req.TenantId,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	return &employeepb.EmployeeResponse{
		Id:        e.ID.String(),
		TenantId:  e.TenantID.String(),
		FirstName: e.FirstName,
		LastName:  e.LastName,
	}, nil
}

func (h *EmployeeHandler) ListEmployee(ctx context.Context, req *employeepb.ListEmployeeRequest) (*employeepb.ListEmployeeResponse, error) {

	list, err := h.service.List(
		ctx,
		req.TenantId,
	)
	if err != nil {
		return nil, err
	}

	res := make([]*employeepb.EmployeeResponse, 0)

	for _, e := range list {

		res = append(res, &employeepb.EmployeeResponse{
			Id:        e.ID.String(),
			TenantId:  e.TenantID.String(),
			FirstName: e.FirstName,
			LastName:  e.LastName,
		})

	}

	return &employeepb.ListEmployeeResponse{
		Employees: res,
	}, nil
}

func (h *EmployeeHandler) DeleteEmployee(ctx context.Context, req *employeepb.DeleteEmployeeRequest) (*employeepb.DeleteEmployeeResponse, error) {

	err := h.service.Delete(
		ctx,
		req.TenantId,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	return &employeepb.DeleteEmployeeResponse{
		Success: true,
	}, nil
}
