package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	authpb "coraflow-erp-api/proto/user/auth/v1"
	userpb "coraflow-erp-api/proto/user/user/v1"

	"coraflow-erp-api/services/user-service/internal/grpc/handler"
)

func Start(port string, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) error {
	server := grpc.NewServer()

	userpb.RegisterUserServiceServer(server, userHandler)
	authpb.RegisterAuthServiceServer(server, authHandler)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	return server.Serve(l)
}
