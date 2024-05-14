package services

import (
	userproto "user-service/genproto/user"
	"user-service/internal/usecase"

	"go.uber.org/zap"
)

type userRPC struct {
	logger             *zap.Logger
	ownerUsecase       usecase.Owner
	workerUsecase      usecase.Worker
	geolocationUsecase usecase.Geolocation
	userproto.UnimplementedUserServiceServer
}
	// CREATE TABLE IF NOT EXISTS categories_products (
	// 	id SERIAL PRIMARY KEY,
	// 	product_id INT REFERENCES products(id),
	// 	category_id INT REFERENCES categories(id), 
	// 	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	// 	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	// );
		
func NewRPC(logger *zap.Logger,
	ownerUsecase usecase.Owner,
	workerUsecase usecase.Worker,
	geolocationUsecase usecase.Geolocation) userproto.UserServiceServer {
	return &userRPC{
		logger:             logger,
		ownerUsecase:       ownerUsecase,
		workerUsecase:      workerUsecase,
		geolocationUsecase: geolocationUsecase,
	}
}
