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

	userRepo := NewOwnersRepo(s.DB)
	ctx := context.Background()

	// struct for create user
	user := entity.Owner{
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
	user.Id = uuid.New().String()

	updOwner := entity.Owner{
		Id:          user.Id,
		FullName:    "updateFullName",
		CompanyName: "updateCompanyName",
		Email:       "updateEmail",
		Password:    "updatePassword",
		Avatar:      "updateAvatar",
		Tax:         12,
	}

	// check create user method
	err = userRepo.CreateOwner(ctx, &user)
	s.Suite.NoError(err)
	Params := make(map[string]string)
	Params["id"] = user.Id

	// check get user method
	getOwner, err := userRepo.GetOwner(ctx, Params)
	s.Suite.NoError(err)
	s.Suite.NotNil(getOwner)
	s.Suite.Equal(getOwner.Id, user.Id)
	s.Suite.Equal(getOwner.FullName, user.FullName)
	s.Suite.Equal(getOwner.CompanyName, user.CompanyName)
	s.Suite.Equal(getOwner.Email, user.Email)
	s.Suite.Equal(getOwner.Password, user.Password)
	s.Suite.Equal(getOwner.Avatar, user.Avatar)
	s.Suite.Equal(getOwner.Tax, user.Tax)

	// check update user method
	err = userRepo.UpdateOwner(ctx, &updOwner)
	s.Suite.NoError(err)
	updGetOwner, err := userRepo.GetOwner(ctx, Params)
	s.Suite.NoError(err)
	s.Suite.NotNil(updGetOwner)
	s.Suite.Equal(updGetOwner.Id, updOwner.Id)
	s.Suite.Equal(updGetOwner.FullName, updOwner.FullName)
	s.Suite.Equal(getOwner.CompanyName, user.CompanyName)
	s.Suite.Equal(getOwner.Email, user.Email)
	s.Suite.Equal(getOwner.Password, user.Password)
	s.Suite.Equal(getOwner.Avatar, user.Avatar)
	s.Suite.Equal(getOwner.Tax, user.Tax)

	// check getAllOwners method
	getAllOwners, err := userRepo.ListOwner(ctx, 5, 1, nil)
	s.Suite.NoError(err)
	s.Suite.NotNil(getAllOwners)


	// ---------------------------------
	// req := entity.CheckFieldReq{
	// 	Value: updOwner.PhoneNumber,
	// 	Field: "phone_number",
	// }

	// check CheckField user method
	// result, err := userRepo.CheckField(ctx, &req)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(updGetOwner)
	// s.Suite.Equal(result.Status, true)

	// check IfExists user method
	// if_exists_req := entity.IfExistsReq{
	// 	PhoneNumber: updOwner.PhoneNumber,
	// }
	// status, err := userRepo.IfExists(ctx, &if_exists_req)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(updGetOwner)
	// s.Suite.Equal(status.IsExistsReq, true)

	// check ChangePassword user method
	// change_password_req := entity.ChangeOwnerPasswordReq{
	// 	PhoneNumber: updOwner.PhoneNumber,
	// 	Password:    "new_password",
	// }

	// resp_change_password, err := userRepo.ChangePassword(ctx, &change_password_req)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(resp_change_password)
	// s.Suite.Equal(resp_change_password.Status, true)

	// check UpdateRefreshToken user method
	// req_update_refresh_token := entity.UpdateRefreshTokenReq{
	// 	OwnerId:       updOwner.Id,
	// 	RefreshToken: "new_refresh_token",
	// }
	// resp_update_refresh_token, err := userRepo.UpdateRefreshToken(ctx, &req_update_refresh_token)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(resp_update_refresh_token)
	// s.Suite.Equal(resp_update_refresh_token.Status, true)

	// check delete user method
	err = userRepo.DeleteOwner(ctx, user.Id)
	s.Suite.NoError(err)

}

func TestOwnerTestSuite(t *testing.T) {
	suite.Run(t, new(OwnerReposisitoryTestSuite))
}
