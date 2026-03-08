package grpc

import (
	"fmt"
	"net"

	tenantpb "coraflow-erp-api/proto/tenant/tenant/v1"
	"coraflow-erp-api/services/tenant-service/internal/grpc/handler"

	"google.golang.org/grpc"
)

func Start(port string, tenantHandler *handler.TenantHandler) error {

	server := grpc.NewServer()

	tenantpb.RegisterTenantServiceServer(server, tenantHandler)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	return server.Serve(l)
}
