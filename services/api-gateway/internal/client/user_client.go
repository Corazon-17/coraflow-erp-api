package client

import (
	authpb "coraflow-erp-api/proto/user/auth/v1"
	"coraflow-erp-api/shared/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAuthClient(cfg *config.Config) (authpb.AuthServiceClient, error) {

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return authpb.NewAuthServiceClient(conn), nil
}
