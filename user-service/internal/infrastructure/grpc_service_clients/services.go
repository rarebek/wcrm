package grpc_service_clients

import (
	// "fmt"
	// postproto "user-service/genproto/post_service"
	"user-service/internal/pkg/config"

	// "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type ServiceClients interface {
	// PostService() postproto.PostServiceClient
	Close()
}

type serviceClients struct {
	// postService    postproto.PostServiceClient
	services       []*grpc.ClientConn
}

func New(config *config.Config) (ServiceClients, error) {
	// connPostService, err := grpc.Dial(
	// 	fmt.Sprintf("%s%s", config.PostService.Host, config.PostService.Port),
	// 	grpc.WithInsecure(),
	// 	// grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	// 	// grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	// )

	// if err != nil {
	// 	return nil, err
	// }


	return &serviceClients{
		// postService:    postproto.NewPostServiceClient(connPostService),
		services:       []*grpc.ClientConn{},
	}, nil
}

func (s *serviceClients) Close() {
	// closing investment service
	for _, conn := range s.services {
		conn.Close()
	}
}

// func (s *serviceClients) PostService() postproto.PostServiceClient {
// 	return s.postService
// }
