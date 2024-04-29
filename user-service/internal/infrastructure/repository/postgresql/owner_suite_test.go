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

type OwnerReposisitoryTestSuite struct {
	suite.Suite
	Config     *config.Config
	DB         *postgres.PostgresDB
	repo       repo.Owners
	ctxTimeout time.Duration
}

func NewOwnerService(ctxTimeout time.Duration, repo repo.Owners, config *config.Config) OwnerReposisitoryTestSuite {
	return OwnerReposisitoryTestSuite{
		Config:     config,
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

// test func
func (s *OwnerReposisitoryTestSuite) TestOwnerCRUD() {

	config := config.New()

	db, err := postgres.New(config)
	if err != nil {
		s.T().Fatal("Error initializing database connection:", err)
	}

	s.DB = db

	ownerRepo := NewOwnersRepo(s.DB)
	ctx := context.Background()

	// struct for create owner
	owner := entity.Owner{
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
	owner.Id = uuid.New().String()

	updOwner := entity.Owner{
		Id:          owner.Id,
		FullName:    "updateFullName",
		CompanyName: "updateCompanyName",
		Email:       "updateEmail",
		Password:    "updatePassword",
		Avatar:      "updateAvatar",
		Tax:         12,
	}

	// check create owner method
	// err = ownerRepo.Create(ctx, &owner)
	s.Suite.NoError(err)
	Params := make(map[string]string)
	Params["id"] = owner.Id

	// check get owner method
	getOwner, err := ownerRepo.Get(ctx, Params)
	s.Suite.NoError(err)
	s.Suite.NotNil(getOwner)
	s.Suite.Equal(getOwner.Id, owner.Id)
	s.Suite.Equal(getOwner.FullName, owner.FullName)
	s.Suite.Equal(getOwner.CompanyName, owner.CompanyName)
	s.Suite.Equal(getOwner.Email, owner.Email)
	s.Suite.Equal(getOwner.Password, owner.Password)
	s.Suite.Equal(getOwner.Avatar, owner.Avatar)
	s.Suite.Equal(getOwner.Tax, owner.Tax)

	// check update owner method
	_, err = ownerRepo.Update(ctx, &updOwner)
	s.Suite.NoError(err)
	updGetOwner, err := ownerRepo.Get(ctx, Params)
	s.Suite.NoError(err)
	s.Suite.NotNil(updGetOwner)
	s.Suite.Equal(updGetOwner.Id, updOwner.Id)
	s.Suite.Equal(updGetOwner.FullName, updOwner.FullName)
	s.Suite.Equal(getOwner.CompanyName, owner.CompanyName)
	s.Suite.Equal(getOwner.Email, owner.Email)
	s.Suite.Equal(getOwner.Password, owner.Password)
	s.Suite.Equal(getOwner.Avatar, owner.Avatar)
	s.Suite.Equal(getOwner.Tax, owner.Tax)

	// check getAllOwners method
	getAllOwners, err := ownerRepo.List(ctx, 5, 1, nil)
	s.Suite.NoError(err)
	s.Suite.NotNil(getAllOwners)

	// check getAllOwners method
	req := entity.CheckFieldRequest{
		Field: "email",
		Value: updOwner.Email,
	}

	// check CheckField owner method
	result, err := ownerRepo.CheckField(ctx, req.Field, req.Value)
	s.Suite.NoError(err)
	s.Suite.NotNil(updGetOwner)
	s.Suite.Equal(result, true)

	// check IfExists owner method
	// if_exists_req := entity.IfExistsReq{
	// 	PhoneNumber: updOwner.PhoneNumber,
	// }
	// status, err := ownerRepo.IfExists(ctx, &if_exists_req)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(updGetOwner)
	// s.Suite.Equal(status.IsExistsReq, true)

	// check ChangePassword owner method
	// change_password_req := entity.ChangeOwnerPasswordReq{
	// 	PhoneNumber: updOwner.PhoneNumber,
	// 	Password:    "new_password",
	// }

	// resp_change_password, err := ownerRepo.ChangePassword(ctx, &change_password_req)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(resp_change_password)
	// s.Suite.Equal(resp_change_password.Status, true)

	// check UpdateRefreshToken owner method
	// req_update_refresh_token := entity.UpdateRefreshTokenReq{
	// 	OwnerId:       updOwner.Id,
	// 	RefreshToken: "new_refresh_token",
	// }
	// resp_update_refresh_token, err := ownerRepo.UpdateRefreshToken(ctx, &req_update_refresh_token)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(resp_update_refresh_token)
	// s.Suite.Equal(resp_update_refresh_token.Status, true)

	// check delete owner method
	_, err = ownerRepo.Delete(ctx, owner.Id)
	s.Suite.NoError(err)

}

func TestOwnerTestSuite(t *testing.T) {
	suite.Run(t, new(OwnerReposisitoryTestSuite))
}
