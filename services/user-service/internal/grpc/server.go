package grpc

import (
	authpb "coraflow-erp-api/proto/user/auth/v1"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer
}

func NewServer() *grpc.Server {

	s := grpc.NewServer()

	authpb.RegisterAuthServiceServer(s, &Server{})

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
