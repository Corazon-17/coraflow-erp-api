package grpc

import (
	"net"

	departmentpb "coraflow-erp-api/proto/hr/department/v1"
	employeepb "coraflow-erp-api/proto/hr/employee/v1"

	"google.golang.org/grpc"
)

type Server struct {
	employeepb.UnimplementedEmployeeServiceServer
	departmentpb.UnimplementedDepartmentServiceServer
}

func NewServer() *grpc.Server {

	s := grpc.NewServer()

	handler := &Server{}

	employeepb.RegisterEmployeeServiceServer(s, handler)
	departmentpb.RegisterDepartmentServiceServer(s, handler)

	return s
}

func Start(port string) error {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	s := NewServer()

	return s.Serve(lis)
}
