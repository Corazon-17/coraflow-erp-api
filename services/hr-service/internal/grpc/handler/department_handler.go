package handler

import (
	"context"

	departmentpb "coraflow-erp-api/proto/hr/department/v1"
	"coraflow-erp-api/services/hr-service/internal/service"
)

type DepartmentHandler struct {
	departmentpb.UnimplementedDepartmentServiceServer
	service *service.DepartmentService
}

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		service: s,
	}
}

func (h *DepartmentHandler) CreateDepartment(ctx context.Context, req *departmentpb.CreateDepartmentRequest) (*departmentpb.DepartmentResponse, error) {

	dept, err := h.service.Create(
		ctx,
		req.TenantId,
		req.Name,
		req.ParentId,
	)
	if err != nil {
		return nil, err
	}

	return &departmentpb.DepartmentResponse{
		Id:       dept.ID.String(),
		TenantId: dept.TenantID.String(),
		Name:     dept.Name,
	}, nil
}

func (h *DepartmentHandler) GetDepartment(ctx context.Context, req *departmentpb.GetDepartmentRequest) (*departmentpb.DepartmentResponse, error) {

	dept, err := h.service.Get(
		ctx,
		req.TenantId,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	return &departmentpb.DepartmentResponse{
		Id:       dept.ID.String(),
		TenantId: dept.TenantID.String(),
		Name:     dept.Name,
	}, nil
}

func (h *DepartmentHandler) ListDepartment(ctx context.Context, req *departmentpb.ListDepartmentRequest) (*departmentpb.ListDepartmentResponse, error) {

	list, err := h.service.List(
		ctx,
		req.TenantId,
	)
	if err != nil {
		return nil, err
	}

	res := make([]*departmentpb.DepartmentResponse, 0)

	for _, d := range list {

		res = append(res, &departmentpb.DepartmentResponse{
			Id:       d.ID.String(),
			TenantId: d.TenantID.String(),
			Name:     d.Name,
		})

	}

	return &departmentpb.ListDepartmentResponse{
		Departments: res,
	}, nil
}

func (h *DepartmentHandler) DeleteDepartment(ctx context.Context, req *departmentpb.DeleteDepartmentRequest,
) (*departmentpb.DeleteDepartmentResponse, error) {

	err := h.service.Delete(
		ctx,
		req.TenantId,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	return &departmentpb.DeleteDepartmentResponse{
		Success: true,
	}, nil
}
