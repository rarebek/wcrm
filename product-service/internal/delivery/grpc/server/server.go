// Package server provides
package server

import (
	"fmt"
	"net"
	"wcrm/product-service/internal/pkg/config"

	"google.golang.org/grpc"
)

func Run(config *config.Config, server *grpc.Server) error {
	lis, err := net.Listen("tcp", config.RPCPort)
	if err != nil {
		return fmt.Errorf("gRPC fatal to listen on %s %w", config.RPCPort, err)
	}

	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}
