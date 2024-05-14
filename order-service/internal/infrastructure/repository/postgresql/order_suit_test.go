package postgresql

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	pb "projects/order-service/genproto/order"
	"projects/order-service/internal/pkg/config"
	db "projects/order-service/internal/pkg/postgres"

	model "projects/order-service/internal/entity"
	Order "projects/order-service/internal/infrastructure/repository"
)

type OrderRepositrySuiteTest struct {
	suite.Suite
	Repository Order.Order
}

func (p *OrderRepositrySuiteTest) SetupSuite() {
	pgPoll, err := db.New(config.New())

	if err != nil {
		log.Fatal("Error while connecting database with suite test")
		return
	}
	p.Repository = NewOrderRepo(pgPoll)
}

// test func
func (p *OrderRepositrySuiteTest) TestProductCRUD() {
	// create comment
	// ---------------------------------------------------------------------------------------------------

	WorkerID := uuid.New().String()
	orderReq := &pb.Order{}
	createdOrder, err := p.Repository.CreateOrder(context.Background(), &model.Order{
		WorkerId:   WorkerID,
		Tax:        orderReq.Tax,
		Discount:   orderReq.Discount,
		TotalPrice: orderReq.TotalPrice,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})

	p.Suite.NoError(err)
	p.Suite.NotNil(createdOrder)
	p.Suite.NotNil(createdOrder.Id)
	p.Suite.NotNil(orderReq.WorkerId)
	p.Suite.Equal(orderReq.Tax, createdOrder.Tax)
	p.Suite.Equal(orderReq.Discount, createdOrder.Discount)
	p.Suite.Equal(orderReq.TotalPrice, createdOrder.TotalPrice)
	p.Suite.NotNil(createdOrder.CreatedAt)
	p.Suite.NotNil(createdOrder.UpdatedAt)

	// ---------------------------------------------------------------------------------------------------------

	// update product
	// ---------------------------------------------------------------------------------------------------------

	WorkerUpdatedOrder := uuid.New().String()

	createdOrder.WorkerId = WorkerUpdatedOrder
	createdOrder.Tax = 125
	createdOrder.Discount = 100
	createdOrder.TotalPrice = 156

	updatedOrder, err := p.Repository.UpdateOrder(context.Background(), createdOrder)
	p.Suite.NoError(err)
	p.Suite.NotNil(updatedOrder)
	p.Suite.NotNil(updatedOrder.Id)
	p.Suite.NotNil(createdOrder.WorkerId)
	p.Suite.Equal(createdOrder.Tax, updatedOrder.Tax)
	p.Suite.Equal(createdOrder.Discount, updatedOrder.Discount)
	p.Suite.Equal(createdOrder.TotalPrice, updatedOrder.TotalPrice)

	// ----------------------------------------------------------------------------------------------------------

	// ----------------------------------------------------------------------------------------------------------
	// get product

	// filter := make(map[string]int64)
	// filter["id"] = int64(updatedOrder.Id)
	// GetOrder, err := p.Repository.GetOrder(context.Background(), filter)
	// p.Suite.NoError(err)
	// p.Suite.NotNil(GetOrder)
	// p.Suite.NotNil(updatedOrder.WorkerId)
	// p.Suite.Equal(updatedOrder.Tax, GetOrder.Tax)
	// p.Suite.Equal(updatedOrder.Discount, GetOrder.Discount)
	// p.Suite.Equal(updatedOrder.TotalPrice, GetOrder.TotalPrice)
	// // ----------------------------------------------------------------------------------------------------------

	// ----------------------------------------------------------------------------------------------------------

	// get all comment
	allComment, err := p.Repository.GetOrders(context.Background(), 1, 5, map[string]string{})
	p.Suite.NoError(err)
	p.Suite.NotNil(allComment)
	// ----------------------------------------------------------------------------------------------------------

}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositrySuiteTest))
}
