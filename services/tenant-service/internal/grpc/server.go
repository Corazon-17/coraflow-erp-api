package grpc

import (
	tenantpb "coraflow-erp-api/proto/tenant/tenant/v1"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	tenantpb.UnimplementedTenantServiceServer
}

func NewServer() *grpc.Server {

	s := grpc.NewServer()

	tenantpb.RegisterTenantServiceServer(s, &Server{})

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
