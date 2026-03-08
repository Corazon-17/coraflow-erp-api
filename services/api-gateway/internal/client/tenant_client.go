package client

import (
	tenantpb "coraflow-erp-api/proto/tenant/tenant/v1"
	"coraflow-erp-api/shared/config"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TenantClient struct {
	Client tenantpb.TenantServiceClient
}

func NewTenantClient(cfg *config.Config) (*TenantClient, error) {

	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.ServerHost, cfg.TenantServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	client := tenantpb.NewTenantServiceClient(conn)

	return &TenantClient{Client: client}, nil
}
