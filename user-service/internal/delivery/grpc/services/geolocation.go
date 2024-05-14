package services

import (
	"context"
	"strconv"
	userproto "user-service/genproto/user"
	"user-service/internal/entity"

	"user-service/internal/pkg/otlp"

	"go.opentelemetry.io/otel/attribute"
)

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

	req, err := s.geolocationUsecase.Delete(ctx, in.Id)

	if err != nil {
		s.logger.Error(err.Error())
		return &userproto.DeletedGeolocation{Status: false}, nil
	}
	return &userproto.DeletedGeolocation{Status: req.Check}, nil
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

	offset := in.Limit * (in.Page - 1)

	geolocations, err := s.geolocationUsecase.List(ctx, in.OwnerId, uint64(in.Limit), uint64(offset), map[string]string{})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	var response userproto.GetAllGeolocationResponse

	for _, u := range geolocations.Geolocations {
		response.AllGeolocations = append(response.AllGeolocations, &userproto.Geolocation{
			Id:        u.Id,
			Latitude:  u.Latitude,
			Longitude: u.Longitude,
			OwnerId:   u.OwnerId,
		})
	}
	response.Count = int64(geolocations.Count)

	return &response, nil
}
