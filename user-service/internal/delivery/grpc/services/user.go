package services

import (
	"context"
	"strconv"
	"time"
	userproto "user-service/genproto/user_service"
	"user-service/internal/entity"
	"user-service/internal/usecase"
	"user-service/internal/usecase/event"

	"user-service/internal/pkg/otlp"

	"go.opentelemetry.io/otel/attribute"

	"go.uber.org/zap"
)

type userRPC struct {
	logger             *zap.Logger
	ownerUsecase       usecase.Owner
	workerUsecase      usecase.Worker
	geolocationUsecase usecase.Geolocation
	brokerProducer     event.BrokerProducer
}

func NewRPC(logger *zap.Logger,
	ownerUsecase usecase.Owner,
	workerUsecase usecase.Worker,
	geolocationUsecase usecase.Geolocation,
	brokerProducer event.BrokerProducer) userproto.UserServiceServer {
	return &userRPC{
		logger:             logger,
		ownerUsecase:       ownerUsecase,
		workerUsecase:      workerUsecase,
		geolocationUsecase: geolocationUsecase,
		brokerProducer:     brokerProducer,
	}
}

func (s userRPC) CreateOwner(ctx context.Context, in *userproto.Owner) (*userproto.Owner, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "CreateOwner")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()
	owner, err := s.ownerUsecase.Create(ctx, &entity.Owner{
		Id:          in.Id,
		FullName:    in.FullName,
		CompanyName: in.CompanyName,
		Email:       in.Email,
		Password:    in.Password,
		Avatar:      in.Avatar,
		Tax:         in.Tax,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Owner{
		Id:          owner.Id,
		FullName:    owner.FullName,
		CompanyName: owner.CompanyName,
		Email:       owner.Email,
		Password:    owner.Password,
		Avatar:      owner.Avatar,
		Tax:         owner.Tax,
		CreatedAt:   owner.CreatedAt.Format("Jan 2, 2006 - 03:04 PM"),
		UpdatedAt:   owner.UpdatedAt.Format("Jan 2, 2006 - 03:04 PM"),
	}, nil
}

func (s userRPC) UpdateOwner(ctx context.Context, in *userproto.Owner) (*userproto.Owner, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdateOwner")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()
	owner, err := s.ownerUsecase.Update(ctx, &entity.Owner{
		Id:          in.Id,
		FullName:    in.FullName,
		CompanyName: in.CompanyName,
		Email:       in.Email,
		Password:    in.Password,
		Avatar:      in.Avatar,
		Tax:         in.Tax,
		UpdatedAt:   time.Now().UTC(),
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Owner{
		Id:          owner.Id,
		FullName:    owner.FullName,
		CompanyName: owner.CompanyName,
		Avatar:      owner.Avatar,
		Email:       owner.Email,
		Password:    owner.Password,
		Tax:         owner.Tax,
		CreatedAt:   in.CreatedAt,
		UpdatedAt:   in.UpdatedAt,
	}, nil
}

func (s userRPC) DeleteOwner(ctx context.Context, in *userproto.GetOwnerRequest) (*userproto.DeletedOwner, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "DeleteOwner")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()
	if err := s.ownerUsecase.Delete(ctx, in.Id); err != nil {
		s.logger.Error(err.Error())
		return &userproto.DeletedOwner{Status: false}, err
	}
	return &userproto.DeletedOwner{Status: true}, nil
}

func (s userRPC) GetOwner(ctx context.Context, in *userproto.GetOwnerRequest) (*userproto.Owner, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetOwner")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()

	user, err := s.ownerUsecase.Get(ctx, map[string]string{
		"id": in.Id,
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Owner{
		Id:          user.Id,
		FullName:    user.FullName,
		CompanyName: user.CompanyName,
		Avatar:      user.Avatar,
		Email:       user.Email,
		Password:    user.Password,
		Tax:         user.Tax,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s userRPC) ListOwner(ctx context.Context, in *userproto.GetAllOwnerRequest) (*userproto.GetAllOwnerResponse, error) {
	//tracing start
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "ListOwner")
	defer span.End()
	//tracing end
	offset := in.Limit * (in.Page - 1)
	users, err := s.ownerUsecase.List(ctx, uint64(in.Limit), uint64(offset), map[string]string{})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response userproto.GetAllOwnerResponse
	for _, u := range users {

		temp := &userproto.Owner{
			Id:          u.Id,
			FullName:    u.FullName,
			CompanyName: u.CompanyName,
			Avatar:      u.Avatar,
			Email:       u.Email,
			Password:    u.Password,
			Tax:         u.Tax,
			CreatedAt:   u.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   u.UpdatedAt.Format(time.RFC3339),
		}

		response.AllOwners = append(response.AllOwners, temp)
	}

	return &response, nil
}

func (s userRPC) CreateWorker(ctx context.Context, in *userproto.Worker) (*userproto.Worker, error) {
	//tracing start
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "CreateWorker")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()
	//tracing end
	worker, err := s.workerUsecase.Create(ctx, &entity.Worker{
		Id:        in.Id,
		FullName:  in.FullName,
		LoginKey:  in.LoginKey,
		Password:  in.Password,
		OwnerId:   in.OwnerId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Worker{
		Id:        worker.Id,
		FullName:  worker.FullName,
		LoginKey:  worker.LoginKey,
		Password:  worker.Password,
		OwnerId:   worker.OwnerId,
		CreatedAt: worker.CreatedAt.Format("Jan 2, 2006 - 03:04 PM"),
		UpdatedAt: worker.UpdatedAt.Format("Jan 2, 2006 - 03:04 PM"),
	}, nil
}

func (s userRPC) UpdateWorker(ctx context.Context, in *userproto.Worker) (*userproto.Worker, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdateWorker")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()
	//tracing end
	worker, err := s.workerUsecase.Update(ctx, &entity.Worker{
		Id:        in.Id,
		FullName:  in.FullName,
		LoginKey:  in.LoginKey,
		Password:  in.Password,
		OwnerId:   in.OwnerId,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &userproto.Worker{
		Id:        worker.Id,
		FullName:  worker.FullName,
		LoginKey:  worker.LoginKey,
		Password:  worker.Password,
		OwnerId:   worker.OwnerId,
		CreatedAt: worker.CreatedAt.Format("Jan 2, 2006 - 03:04 PM"),
		UpdatedAt: worker.UpdatedAt.Format("Jan 2, 2006 - 03:04 PM"),
	}, nil
}

func (s userRPC) DeleteWorker(ctx context.Context, in *userproto.GetWorkerRequest) (*userproto.DeletedWorker, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "DeleteWorker")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()
	//tracing end
	if err := s.workerUsecase.Delete(ctx, in.Id); err != nil {
		s.logger.Error(err.Error())
		return &userproto.DeletedWorker{Status: false}, err
	}
	return &userproto.DeletedWorker{Status: true}, nil
}

func (s userRPC) GetWorker(ctx context.Context, in *userproto.GetWorkerRequest) (*userproto.Worker, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetWorker")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()
	//tracing end
	user, err := s.workerUsecase.Get(ctx, map[string]string{
		"id": in.Id,
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Worker{
		Id:        user.Id,
		FullName:  user.FullName,
		LoginKey:  user.LoginKey,
		Password:  user.Password,
		OwnerId:   user.OwnerId,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s userRPC) ListWorker(ctx context.Context, in *userproto.GetAllWorkerRequest) (*userproto.GetAllWorkerResponse, error) {
	//tracing start
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "ListWorker")
	defer span.End()
	//tracing end

	offset := in.Limit * (in.Page - 1)
	users, err := s.workerUsecase.List(ctx, uint64(in.Limit), uint64(offset), map[string]string{})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response userproto.GetAllWorkerResponse
	for _, u := range users {

		temp := &userproto.Worker{
			Id:        u.Id,
			FullName:  u.FullName,
			LoginKey:  u.LoginKey,
			Password:  u.Password,
			OwnerId:   u.OwnerId,
			CreatedAt: u.CreatedAt.Format(time.RFC3339),
			UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
		}

		response.AllWorkers = append(response.AllWorkers, temp)
	}

	return &response, nil
}

// --------------------------------------

func (s userRPC) CreateGeolocation(ctx context.Context, in *userproto.Geolocation) (*userproto.Geolocation, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "CreateGeolocation")
	span.SetAttributes(
		attribute.Key("id").String(strconv.Itoa(int(in.Id))),
	)
	defer span.End()
	//tracing end
	geolocation, err := s.geolocationUsecase.Create(ctx, &entity.Geolocation{
		Id:        in.Id,
		Latitude:  in.Latitude,
		Longitude: in.Longitude,
		OwnerId:   in.OwnerId,
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Geolocation{
		Id:        geolocation.Id,
		Latitude:  geolocation.Latitude,
		Longitude: geolocation.Longitude,
		OwnerId:   geolocation.OwnerId,
	}, nil
}

func (s userRPC) UpdateGeolocation(ctx context.Context, in *userproto.Geolocation) (*userproto.Geolocation, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "UpdateGeolocation")
	span.SetAttributes(
		attribute.Key("id").String(strconv.Itoa(int(in.Id))),
	)
	defer span.End()
	geolocation, err := s.geolocationUsecase.Update(ctx, &entity.Geolocation{
		Id:        in.Id,
		Latitude:  in.Latitude,
		Longitude: in.Longitude,
		OwnerId:   in.OwnerId,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &userproto.Geolocation{
		Id:        geolocation.Id,
		Latitude:  geolocation.Latitude,
		Longitude: geolocation.Longitude,
		OwnerId:   geolocation.OwnerId,
	}, nil
}

func (s userRPC) DeleteGeolocation(ctx context.Context, in *userproto.GetGeolocationRequest) (*userproto.DeletedGeolocation, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "DeleteGeolocation")
	span.SetAttributes(
		attribute.Key("id").String(strconv.Itoa(int(in.Id))),
	)
	defer span.End()
	//tracing end
	if err := s.geolocationUsecase.Delete(ctx, in.Id); err != nil {
		s.logger.Error(err.Error())
		return &userproto.DeletedGeolocation{Status: false}, err
	}
	return &userproto.DeletedGeolocation{Status: true}, nil
}

func (s userRPC) GetGeolocation(ctx context.Context, in *userproto.GetGeolocationRequest) (*userproto.Geolocation, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetGeolocation")
	span.SetAttributes(
		attribute.Key("id").String(strconv.Itoa(int(in.Id))),
	)
	defer span.End()
	//tracing end
	user, err := s.geolocationUsecase.Get(ctx, map[string]int64{
		"id": in.Id,
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Geolocation{
		Id:        user.Id,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
		OwnerId:   user.OwnerId,
	}, nil
}

func (s userRPC) ListGeolocation(ctx context.Context, in *userproto.GetAllGeolocationRequest) (*userproto.GetAllGeolocationResponse, error) {
	//tracing start
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "ListGeolocation")
	defer span.End()
	//tracing end
	users, err := s.geolocationUsecase.List(ctx, map[string]string{})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response userproto.GetAllGeolocationResponse
	for _, u := range users {

		temp := &userproto.Geolocation{
			Id:        u.Id,
			Latitude:  u.Latitude,
			Longitude: u.Longitude,
			OwnerId:   u.OwnerId,
		}

		response.AllGeolocations = append(response.AllGeolocations, temp)
	}

	return &response, nil
}

func (s userRPC) CheckFieldOwner(ctx context.Context, in *userproto.CheckFieldRequest) (*userproto.CheckFieldResponse, error) {
	//tracing start
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "CheckFieldOwner")
	span.SetAttributes(
		attribute.Key(in.Field).String(in.Value),
	)
	defer span.End()
	//tracing end

	exist, err := s.ownerUsecase.CheckField(ctx, in.Field, in.Value)
	if err != nil {
		s.logger.Error(err.Error())
		return &userproto.CheckFieldResponse{Exist: exist}, err
	}

	return &userproto.CheckFieldResponse{Exist: exist}, nil
}

func (s userRPC) CheckFieldWorker(ctx context.Context, in *userproto.CheckFieldRequest) (*userproto.CheckFieldResponse, error) {
	//tracing start
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "CheckFieldWorker")
	span.SetAttributes(
		attribute.Key(in.Field).String(in.Value),
	)
	defer span.End()
	//tracing end

	exist, err := s.workerUsecase.CheckField(ctx, in.Field, in.Value)
	if err != nil {
		s.logger.Error(err.Error())
		return &userproto.CheckFieldResponse{Exist: exist}, err
	}

	return &userproto.CheckFieldResponse{Exist: exist}, nil
}
