package postgresql

import (
	"context"
	"testing"
	"time"
	"user-service/internal/entity"
	repo "user-service/internal/infrastructure/repository"
	"user-service/internal/pkg/config"
	"user-service/internal/pkg/postgres"

	"github.com/google/uuid"

	"github.com/stretchr/testify/suite"
)

type GeolocationReposisitoryTestSuite struct {
	suite.Suite
	Config     *config.Config
	DB         *postgres.PostgresDB
	repo       repo.Geolocations
	ctxTimeout time.Duration
}

func NewGeolocationService(ctxTimeout time.Duration, repo repo.Geolocations, config *config.Config) GeolocationReposisitoryTestSuite {
	return GeolocationReposisitoryTestSuite{
		Config:     config,
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

// test func
func (s *GeolocationReposisitoryTestSuite) TestGeolocationCRUD() {

	config := config.New()

	db, err := postgres.New(config)
	if err != nil {
		s.T().Fatal("Error initializing database connection:", err)
	}

	s.DB = db

	userRepo := NewGeolocationsRepo(s.DB)
	ownerRepo := NewOwnersRepo(s.DB)
	ctx := context.Background()

	owner := entity.Owner{
		Id:          uuid.New().String(),
		FullName:    "testFullName",
		CompanyName: "testCompanyName",
		Email:       "testEmail",
		Password:    "testPassword",
		Avatar:      "testAvatar",
		Tax:         12,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
	// uuid generating

	err = ownerRepo.Create(ctx, &owner)
	s.Suite.NoError(err)


	geolocation := entity.Geolocation{
		Latitude:  40.7128,
		Longitude: -74.0060,
		OwnerId:   owner.Id,
	}
	// uuid generating

	updGeolocation := entity.Geolocation{
		Id:        geolocation.Id,
		Latitude:   34.0522,
		Longitude: -118.2437,
		OwnerId:   geolocation.OwnerId,
	}

	// check create geolocation method
	err = userRepo.Create(ctx, &geolocation)
	s.Suite.NoError(err)
	Params := make(map[string]int64)
	Params["id"] = geolocation.Id

	// check get geolocation method
	getGeolocation, err := userRepo.Get(ctx, Params)
	s.Suite.NoError(err)
	s.Suite.NotNil(getGeolocation)
	s.Suite.Equal(getGeolocation.Id, geolocation.Id)
	s.Suite.Equal(getGeolocation.Latitude, geolocation.Latitude)
	s.Suite.Equal(getGeolocation.Longitude, geolocation.Longitude)
	s.Suite.Equal(getGeolocation.OwnerId, geolocation.OwnerId)

	// check update geolocation method
	err = userRepo.Update(ctx, &updGeolocation)
	s.Suite.NoError(err)
	updGetGeolocation, err := userRepo.Get(ctx, Params)
	s.Suite.NoError(err)
	s.Suite.NotNil(updGetGeolocation)
	s.Suite.Equal(updGetGeolocation.Id, updGeolocation.Id)
	s.Suite.Equal(updGetGeolocation.Latitude, updGeolocation.Latitude)
	s.Suite.Equal(getGeolocation.Longitude, geolocation.Longitude)
	s.Suite.Equal(getGeolocation.OwnerId, geolocation.OwnerId)

	// check getAllGeolocations method
	getAllGeolocations, err := userRepo.List(ctx, nil)
	s.Suite.NoError(err)
	s.Suite.NotNil(getAllGeolocations)

	// check delete geolocation method
	err = userRepo.Delete(ctx, geolocation.Id)
	s.Suite.NoError(err)

	err = ownerRepo.Delete(ctx, owner.Id)
	s.Suite.NoError(err)
}

func TestGeolocationTestSuite(t *testing.T) {
	suite.Run(t, new(GeolocationReposisitoryTestSuite))
}
