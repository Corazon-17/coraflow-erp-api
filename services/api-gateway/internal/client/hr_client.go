package client

import (
	departmentpb "coraflow-erp-api/proto/hr/department/v1"
	employeepb "coraflow-erp-api/proto/hr/employee/v1"
	"coraflow-erp-api/shared/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HRClient struct {
	Department departmentpb.DepartmentServiceClient
	Employee   employeepb.EmployeeServiceClient
}

func NewHRClient(cfg *config.Config) (*HRClient, error) {

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.TenantServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &HRClient{
		Department: departmentpb.NewDepartmentServiceClient(conn),
		Employee:   employeepb.NewEmployeeServiceClient(conn),
	}, nil
}
