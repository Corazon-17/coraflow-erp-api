package grpc

import (
	"fmt"
	"net"

	departmentpb "coraflow-erp-api/proto/hr/department/v1"
	employeepb "coraflow-erp-api/proto/hr/employee/v1"
	"coraflow-erp-api/services/hr-service/internal/grpc/handler"

	"google.golang.org/grpc"
)

func Start(port string, departmentHandler *handler.DepartmentHandler, employeeHandler *handler.EmployeeHandler) error {

	server := grpc.NewServer()

	departmentpb.RegisterDepartmentServiceServer(server, departmentHandler)
	employeepb.RegisterEmployeeServiceServer(server, employeeHandler)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	return server.Serve(l)
}
