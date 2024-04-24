package grpc_service_clients

import (
	"fmt"
	pbo "wcrm/product-service/genproto/order"
	"wcrm/product-service/internal/pkg/config"

	"google.golang.org/grpc"
)

type ServiceClients interface {
	OrderService() pbo.OrderServiceClient
	Close()
}

type serviceClients struct {
	services     []*grpc.ClientConn
	orderService pbo.OrderServiceClient
}

func New(config *config.Config) (ServiceClients, error) {

	// dail to order-service
	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", config.OrderService.Host, config.OrderService.Port),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", config.OrderService.Host, config.OrderService.Port)
	}

	return &serviceClients{
		services: []*grpc.ClientConn{},
		orderService: pbo.NewOrderServiceClient(connOrder),
	}, nil
}

func (s *serviceClients) Close() {
	// closing investment service
	for _, conn := range s.services {
		conn.Close()
	}
}

func (s *serviceClients) OrderService() pbo.OrderServiceClient {
	return s.orderService
}
