package grpc_service_clients

import (
	"fmt"

	pbo "api-gateway/genproto/order"
	pbp "api-gateway/genproto/product"
	pbu "api-gateway/genproto/user"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"api-gateway/internal/pkg/config"
)

type ServiceClient interface {
	OrderService() pbo.OrderServiceClient
	ProductService() pbp.ProductServiceClient
	UserService() pbu.UserServiceClient
	Close()
}

type serviceClient struct {
	connections    []*grpc.ClientConn
	userService    pbu.UserServiceClient
	productService pbp.ProductServiceClient
	orderService   pbo.OrderServiceClient
}

func New(cfg *config.Config) (ServiceClient, error) {
	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserService.Host, cfg.UserService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	connProductService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.ProductService.Host, cfg.ProductService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	connOrderService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.OrderService.Host, cfg.OrderService.Port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
	if err != nil {
		return nil, err
	}

	return &serviceClient{
		userService:    pbu.NewUserServiceClient(connUserService),
		productService: pbp.NewProductServiceClient(connProductService),
		orderService:   pbo.NewOrderServiceClient(connOrderService),
		connections: []*grpc.ClientConn{
			connProductService,
			connUserService,
			connOrderService,
		},
	}, nil
}

func (s *serviceClient) OrderService() pbo.OrderServiceClient {
	return s.orderService
}

func (s *serviceClient) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceClient) ProductService() pbp.ProductServiceClient {
	return s.productService
}

func (s *serviceClient) Close() {
	for _, conn := range s.connections {
		if err := conn.Close(); err != nil {
			// should be replaced by logger soon
			fmt.Printf("error while closing grpc connection: %v", err)
		}
	}
}
