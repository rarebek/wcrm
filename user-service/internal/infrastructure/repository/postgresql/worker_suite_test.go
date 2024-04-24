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

type WorkerReposisitoryTestSuite struct {
	suite.Suite
	Config     *config.Config
	DB         *postgres.PostgresDB
	repo       repo.Workers
	ctxTimeout time.Duration
}

func NewWorkerService(ctxTimeout time.Duration, repo repo.Workers, config *config.Config) WorkerReposisitoryTestSuite {
	return WorkerReposisitoryTestSuite{
		Config:     config,
		ctxTimeout: ctxTimeout,
		repo:       repo,
	}
}

// test func
func (s *WorkerReposisitoryTestSuite) TestWorkerCRUD() {

	config := config.New()

	db, err := postgres.New(config)
	if err != nil {
		s.T().Fatal("Error initializing database connection:", err)
	}

	s.DB = db

	userRepo := NewWorkersRepo(s.DB)
	ctx := context.Background()

	// struct for create user
	user := entity.Worker{
		FullName:  "testFullName",
		LoginKey:  "testLoginKey",
		Password:  "testPassword",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	// uuid generating
	user.Id = uuid.New().String()
	user.OwnerId = uuid.New().String()

	updWorker := entity.Worker{
		Id:       user.Id,
		FullName: "updateFullName",
		LoginKey: "updateLoginKey",
		Password: "updatePassword",
		OwnerId:  user.OwnerId,
	}

	// check create user method
	err = userRepo.CreateWorker(ctx, &user)
	s.Suite.NoError(err)
	Params := make(map[string]string)
	Params["id"] = user.Id

	// check get user method
	getWorker, err := userRepo.GetWorker(ctx, Params)
	s.Suite.NoError(err)
	s.Suite.NotNil(getWorker)
	s.Suite.Equal(getWorker.Id, user.Id)
	s.Suite.Equal(getWorker.FullName, user.FullName)
	s.Suite.Equal(getWorker.LoginKey, user.LoginKey)
	s.Suite.Equal(getWorker.Password, user.Password)
	s.Suite.Equal(getWorker.OwnerId, user.OwnerId)

	// check update user method
	err = userRepo.UpdateWorker(ctx, &updWorker)
	s.Suite.NoError(err)
	updGetWorker, err := userRepo.GetWorker(ctx, Params)
	s.Suite.NoError(err)
	s.Suite.NotNil(updGetWorker)
	s.Suite.Equal(updGetWorker.Id, updWorker.Id)
	s.Suite.Equal(updGetWorker.FullName, updWorker.FullName)
	s.Suite.Equal(getWorker.LoginKey, user.LoginKey)
	s.Suite.Equal(getWorker.Password, user.Password)
	s.Suite.Equal(getWorker.OwnerId, user.OwnerId)

	// check getAllWorkers method
	getAllWorkers, err := userRepo.ListWorker(ctx, 5, 1, nil)
	s.Suite.NoError(err)
	s.Suite.NotNil(getAllWorkers)

	// ---------------------------------
	// req := entity.CheckFieldReq{
	// 	Value: updWorker.PhoneNumber,
	// 	Field: "phone_number",
	// }

	// check CheckField user method
	// result, err := userRepo.CheckField(ctx, &req)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(updGetWorker)
	// s.Suite.Equal(result.Status, true)

	// check IfExists user method
	// if_exists_req := entity.IfExistsReq{
	// 	PhoneNumber: updWorker.PhoneNumber,
	// }
	// status, err := userRepo.IfExists(ctx, &if_exists_req)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(updGetWorker)
	// s.Suite.Equal(status.IsExistsReq, true)

	// check ChangePassword user method
	// change_password_req := entity.ChangeWorkerPasswordReq{
	// 	PhoneNumber: updWorker.PhoneNumber,
	// 	Password:    "new_password",
	// }

	// resp_change_password, err := userRepo.ChangePassword(ctx, &change_password_req)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(resp_change_password)
	// s.Suite.Equal(resp_change_password.Status, true)

	// check UpdateRefreshToken user method
	// req_update_refresh_token := entity.UpdateRefreshTokenReq{
	// 	WorkerId:       updWorker.Id,
	// 	RefreshToken: "new_refresh_token",
	// }
	// resp_update_refresh_token, err := userRepo.UpdateRefreshToken(ctx, &req_update_refresh_token)
	// s.Suite.NoError(err)
	// s.Suite.NotNil(resp_update_refresh_token)
	// s.Suite.Equal(resp_update_refresh_token.Status, true)

	// check delete user method
	err = userRepo.DeleteWorker(ctx, user.Id)
	s.Suite.NoError(err)

}

func TestWorkerTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerReposisitoryTestSuite))
}
