package handler

import (
	"context"

	hrpb "coraflow-erp-api/proto/hr/department/v1"
	"coraflow-erp-api/services/hr-service/internal/service"
)

type DepartmentHandler struct {
	hrpb.UnimplementedDepartmentServiceServer
	service *service.DepartmentService
}

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{
		service: s,
	}
}

func (h *DepartmentHandler) CreateDepartment(ctx context.Context, req *hrpb.CreateDepartmentRequest) (*hrpb.DepartmentResponse, error) {

	dept, err := h.service.Create(
		ctx,
		req.TenantId,
		req.Name,
		req.ParentId,
	)
	if err != nil {
		return nil, err
	}

	return &hrpb.DepartmentResponse{
		Id:       dept.ID.String(),
		TenantId: dept.TenantID.String(),
		Name:     dept.Name,
	}, nil
}

func (h *DepartmentHandler) GetDepartment(ctx context.Context, req *hrpb.GetDepartmentRequest) (*hrpb.DepartmentResponse, error) {

	dept, err := h.service.Get(
		ctx,
		req.TenantId,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	return &hrpb.DepartmentResponse{
		Id:       dept.ID.String(),
		TenantId: dept.TenantID.String(),
		Name:     dept.Name,
	}, nil
}

func (h *DepartmentHandler) ListDepartment(ctx context.Context, req *hrpb.ListDepartmentRequest) (*hrpb.ListDepartmentResponse, error) {

	list, err := h.service.List(
		ctx,
		req.TenantId,
	)
	if err != nil {
		return nil, err
	}

	res := make([]*hrpb.DepartmentResponse, 0)

	for _, d := range list {

		res = append(res, &hrpb.DepartmentResponse{
			Id:       d.ID.String(),
			TenantId: d.TenantID.String(),
			Name:     d.Name,
		})

	}

	return &hrpb.ListDepartmentResponse{
		Departments: res,
	}, nil
}

func (h *DepartmentHandler) DeleteDepartment(ctx context.Context, req *hrpb.DeleteDepartmentRequest,
) (*hrpb.DeleteDepartmentResponse, error) {

	err := h.service.Delete(
		ctx,
		req.TenantId,
		req.Id,
	)
	if err != nil {
		return nil, err
	}

	return &hrpb.DeleteDepartmentResponse{
		Success: true,
	}, nil
}
