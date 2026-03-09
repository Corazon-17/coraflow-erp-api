package client

import (
	authpb "coraflow-erp-api/proto/user/auth/v1"
	userpb "coraflow-erp-api/proto/user/user/v1"
	"coraflow-erp-api/shared/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	Auth authpb.AuthServiceClient
	User userpb.UserServiceClient
}

func NewUserClient(cfg *config.Config) (*UserClient, error) {

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &UserClient{
		User: userpb.NewUserServiceClient(conn),
		Auth: authpb.NewAuthServiceClient(conn),
	}, nil
}
