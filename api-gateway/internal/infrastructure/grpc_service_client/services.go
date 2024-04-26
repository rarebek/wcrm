package grpc_service_clients

import (
	"fmt"

	pbu "evrone_service/api_gateway/genproto/user_service"
	pbp "evrone_service/api_gateway/genproto/product"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"

	"evrone_service/api_gateway/internal/pkg/config"
)

type ServiceClient interface {
	ProductService() pbp.ProductServiceClient
	UserService() pbu.UserServiceClient
	Close()
}

type serviceClient struct {
	connections []*grpc.ClientConn
	userService pbu.UserServiceClient
	productService pbp.ProductServiceClient
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



	return &serviceClient{
		userService: pbu.NewUserServiceClient(connUserService),
		productService: pbp.NewProductServiceClient(connProductService),
		connections: []*grpc.ClientConn{
			connProductService,
			connUserService,
		},
	}, nil
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
