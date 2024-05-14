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
