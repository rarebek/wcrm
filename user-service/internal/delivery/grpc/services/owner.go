package services

import (
	"context"
	"time"
	userproto "user-service/genproto/user"
	"user-service/internal/entity"

	"user-service/internal/pkg/otlp"

	"go.opentelemetry.io/otel/attribute"
)

func (s userRPC) CreateOwner(ctx context.Context, in *userproto.Owner) (*userproto.Owner, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "CreateOwner")
	span.SetAttributes(
		attribute.Key("id").String(in.Id),
	)
	defer span.End()
	owner, err := s.ownerUsecase.Create(ctx, &entity.Owner{
		Id:           in.Id,
		FullName:     in.FullName,
		CompanyName:  in.CompanyName,
		Email:        in.Email,
		Password:     in.Password,
		Avatar:       in.Avatar,
		Tax:          in.Tax,
		RefreshToken: in.RefreshToken,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	})

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Owner{
		Id:           owner.Id,
		FullName:     owner.FullName,
		CompanyName:  owner.CompanyName,
		Email:        owner.Email,
		Password:     owner.Password,
		Avatar:       owner.Avatar,
		Tax:          owner.Tax,
		RefreshToken: owner.RefreshToken,
		CreatedAt:    owner.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    owner.UpdatedAt.Format("2006-01-02 15:04:05"),
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
		Id:           in.Id,
		FullName:     in.FullName,
		CompanyName:  in.CompanyName,
		Email:        in.Email,
		Password:     in.Password,
		Avatar:       in.Avatar,
		Tax:          in.Tax,
		RefreshToken: in.RefreshToken,
		UpdatedAt:    time.Now().UTC(),
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Owner{
		Id:           owner.Id,
		FullName:     owner.FullName,
		CompanyName:  owner.CompanyName,
		Avatar:       owner.Avatar,
		Email:        owner.Email,
		Password:     owner.Password,
		Tax:          owner.Tax,
		RefreshToken: owner.RefreshToken,
		CreatedAt:    owner.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    owner.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s userRPC) DeleteOwner(ctx context.Context, in *userproto.IdRequest) (*userproto.DeletedOwner, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "DeleteOwner")
	span.SetAttributes(
		attribute.Key("id").String("delete owner"),
	)
	defer span.End()

	delreq, err := s.ownerUsecase.Delete(ctx, in.Id)

	if err != nil {
		s.logger.Error(err.Error())
		return &userproto.DeletedOwner{Status: false}, nil
	}
	return &userproto.DeletedOwner{Status: delreq.Check}, nil
}

func (s userRPC) GetOwner(ctx context.Context, in *userproto.GetOwnerRequest) (*userproto.Owner, error) {
	//tracing
	ctx, span := otlp.Start(ctx, "user_grpc-delivery", "GetOwner")
	span.SetAttributes(
		attribute.Key("id").String("get owner"),
	)
	defer span.End()

	user, err := s.ownerUsecase.Get(ctx, in.Filter)

	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &userproto.Owner{
		Id:           user.Id,
		FullName:     user.FullName,
		CompanyName:  user.CompanyName,
		Avatar:       user.Avatar,
		Email:        user.Email,
		Password:     user.Password,
		Tax:          user.Tax,
		RefreshToken: user.RefreshToken,
		CreatedAt:    user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    user.UpdatedAt.Format("2006-01-02 15:04:05"),
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

	for _, in := range users.Owners {
		response.Owners = append(response.Owners, &userproto.Owner{
			Id:           in.Id,
			FullName:     in.FullName,
			CompanyName:  in.CompanyName,
			Email:        in.Email,
			Password:     in.Password,
			Avatar:       in.Avatar,
			Tax:          in.Tax,
			RefreshToken: in.RefreshToken,
			CreatedAt:    in.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    in.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	response.Count = int64(users.Count)

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
