package postgresql

// import (
// 	"context"
// 	"testing"
// 	"time"
// 	"user-service/internal/entity"
// 	repo "user-service/internal/infrastructure/repository"
// 	trepo "user-service/internal/infrastructure/repository/postgresql"
// 	"user-service/internal/pkg/config"
// 	"user-service/internal/pkg/postgres"

// 	"github.com/google/uuid"

// 	"github.com/stretchr/testify/suite"
// )

// type WorkerReposisitoryTestSuite struct {
// 	suite.Suite
// 	Config     *config.Config
// 	DB         *postgres.PostgresDB
// 	repo       repo.Workers
// 	ctxTimeout time.Duration
// }

// func NewWorkerService(ctxTimeout time.Duration, repo repo.Workers, config *config.Config) WorkerReposisitoryTestSuite {
// 	return WorkerReposisitoryTestSuite{
// 		Config:     config,
// 		ctxTimeout: ctxTimeout,
// 		repo:       repo,
// 	}
// }

// // test func
// func (s *WorkerReposisitoryTestSuite) TestWorkerCRUD() {

// 	config := config.New()

// 	db, err := postgres.New(config)
// 	if err != nil {
// 		s.T().Fatal("Error initializing database connection:", err)
// 	}

// 	s.DB = db

// 	workerRepo := trepo.NewWorkersRepo(s.DB)
// 	ownerRepo := trepo.NewOwnersRepo(s.DB)
// 	ctx := context.Background()

// 	// struct for create owner
// 	owner := entity.Owner{
// 		Id:          uuid.New().String(),
// 		FullName:    "testFullName",
// 		CompanyName: "testCompanyName",
// 		Email:       "testEmail",
// 		Password:    "testPassword",
// 		Avatar:      "testAvatar",
// 		Tax:         12,
// 		CreatedAt:   time.Now().UTC(),
// 		UpdatedAt:   time.Now().UTC(),
// 	}
// 	// uuid generating

// 	_, err = ownerRepo.Create(ctx, &owner)
// 	s.Suite.NoError(err)

// 	// struct for create worker
// 	worker := entity.Worker{
// 		FullName:  "testFullName",
// 		LoginKey:  "testLoginKey",
// 		Password:  "testPassword",
// 		OwnerId:   owner.Id,
// 		CreatedAt: time.Now().UTC(),
// 		UpdatedAt: time.Now().UTC(),
// 	}
// 	// uuid generating
// 	worker.Id = uuid.New().String()

// 	updWorker := entity.Worker{
// 		Id:       worker.Id,
// 		FullName: "updateFullName",
// 		LoginKey: "updateLoginKey",
// 		Password: "updatePassword",
// 		OwnerId:  owner.Id,
// 	}

// 	// check create worker method
// 	_, err = workerRepo.Create(ctx, &worker)
// 	s.Suite.NoError(err)
// 	Params := make(map[string]string)
// 	Params["id"] = worker.Id

// 	// check get worker method
// 	getWorker, err := workerRepo.Get(ctx, Params)
// 	s.Suite.NoError(err)
// 	s.Suite.NotNil(getWorker)
// 	s.Suite.Equal(getWorker.Id, worker.Id)
// 	s.Suite.Equal(getWorker.FullName, worker.FullName)
// 	s.Suite.Equal(getWorker.LoginKey, worker.LoginKey)
// 	s.Suite.Equal(getWorker.Password, worker.Password)
// 	s.Suite.Equal(getWorker.OwnerId, worker.OwnerId)

// 	// check update worker method
// 	_, err = workerRepo.Update(ctx, &updWorker)
// 	s.Suite.NoError(err)
// 	updGetWorker, err := workerRepo.Get(ctx, Params)
// 	s.Suite.NoError(err)
// 	s.Suite.NotNil(updGetWorker)
// 	s.Suite.Equal(updGetWorker.Id, updWorker.Id)
// 	s.Suite.Equal(updGetWorker.FullName, updWorker.FullName)
// 	s.Suite.Equal(getWorker.LoginKey, worker.LoginKey)
// 	s.Suite.Equal(getWorker.Password, worker.Password)
// 	s.Suite.Equal(getWorker.OwnerId, worker.OwnerId)

// 	// check getAllWorkers method
// 	getAllWorkers, err := workerRepo.List(ctx, 5, 1, nil)
// 	s.Suite.NoError(err)
// 	s.Suite.NotNil(getAllWorkers)

// 	req := entity.CheckFieldRequest{
// 		Field: "login_key",
// 		Value: updWorker.LoginKey,
// 	}

// 	// check CheckField owner method
// 	result, err := workerRepo.CheckField(ctx, req.Field, req.Value)
// 	s.Suite.NoError(err)
// 	s.Suite.Equal(result, true)

// 	// check delete worker method
// 	// err = workerRepo.Delete(ctx, worker.Id)
// 	s.Suite.NoError(err)

// 	// err = ownerRepo.Delete(ctx, owner.Id)
// 	s.Suite.NoError(err)
// }

// func TestWorkerTestSuite(t *testing.T) {
// 	suite.Run(t, new(WorkerReposisitoryTestSuite))
// }
