package services

import (
	"context"
	"time"
	userproto "user-service/genproto/user"
	"user-service/internal/entity"

	"user-service/internal/pkg/otlp"

	"github.com/k0kubun/pp"
	"go.opentelemetry.io/otel/attribute"
)

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
		CreatedAt: worker.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: worker.UpdatedAt.Format("2006-01-02 15:04:05"),
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
		CreatedAt: worker.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: worker.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s userRPC) DeleteWorker(ctx context.Context, in *userproto.IdRequest) (*userproto.DeletedWorker, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "DeleteWorker")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()
	//tracing end

	req, err := s.workerUsecase.Delete(ctx, in.Id)

	if err != nil {
		s.logger.Error(err.Error())
		return &userproto.DeletedWorker{Status: false}, nil
	}
	return &userproto.DeletedWorker{Status: req.Check}, nil
}

func (s userRPC) GetWorker(ctx context.Context, in *userproto.GetWorkerRequest) (*userproto.Worker, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetWorker")
	span.SetAttributes(
		attribute.Key("id").String(in.Filter["id"]),
	)
	defer span.End()
	//tracing end
	user, err := s.workerUsecase.Get(ctx, in.Filter)

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
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s userRPC) ListWorker(ctx context.Context, in *userproto.GetAllWorkerRequest) (*userproto.GetAllWorkerResponse, error) {
	//tracing start
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "ListWorker")
	defer span.End()
	//tracing end
	pp.Println("OWNER ID: ", in.Filter["owner_id"])

	offset := in.Limit * (in.Page - 1)
	workers, err := s.workerUsecase.List(ctx, uint64(in.Limit), uint64(offset), in.Filter)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response userproto.GetAllWorkerResponse
	for _, u := range workers.Workers {
		response.AllWorkers = append(response.AllWorkers, &userproto.Worker{
			Id:        u.Id,
			FullName:  u.FullName,
			LoginKey:  u.LoginKey,
			Password:  u.Password,
			OwnerId:   u.OwnerId,
			CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	response.Count = int64(workers.Count)

	return &response, nil
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
