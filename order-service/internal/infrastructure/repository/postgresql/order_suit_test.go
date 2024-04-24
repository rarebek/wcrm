package postgresql

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"

	pbp "order-service/genproto/order"
	"order-service/internal/pkg/config"
	db "order-service/internal/pkg/postgres"

	model "order-service/internal/entity"
	Order "order-service/internal/infrastructure/repository"
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

// ! SUIT TEST
func (p *OrderRepositrySuiteTest) TestOrderCRUD() {
	// ! CREATE ORDER
	// ---------------------------------------------------------------------------------------------------
	orderReq := &pbp.Order{
		Tax:        2,
		Discount:   15,
		TotalPrice: 365,
	}

	createOrder, err := p.Repository.CreateOrder(context.Background(), &model.Order{
		Tax:        orderReq.Tax,
		Discount:   orderReq.Discount,
		TotalPrice: orderReq.TotalPrice,
	})

	p.Suite.NoError(err)
	p.Suite.NotNil(createOrder)
	p.Suite.NotNil(createOrder.Id)
	p.Suite.Equal(orderReq.WorkerId, createOrder.WorkerId)
	p.Suite.Equal(orderReq.Tax, createOrder.Tax)
	p.Suite.Equal(orderReq.Discount, createOrder.Discount)
	p.Suite.Equal(orderReq.TotalPrice, createOrder.TotalPrice)
	p.Suite.NotNil(createOrder.CreatedAt)
	p.Suite.NotNil(createOrder.UpdatedAt)

	// ---------------------------------------------------------------------------------------------------------

	//! update comment
	// ---------------------------------------------------------------------------------------------------------

	createOrder.WorkerId = "3287f50f-9385-47e4-a910-64d7a1718727"
	createOrder.Tax = 15
	createOrder.Discount = 10000
	createOrder.TotalPrice = 10

	updatedOrder, err := p.Repository.UpdateOrder(context.Background(), createOrder)
	p.Suite.NoError(err)
	p.Suite.NotNil(updatedOrder)
	p.Suite.NotNil(updatedOrder.Id)
	p.Suite.Equal(createOrder.WorkerId, updatedOrder.WorkerId)
	p.Suite.Equal(createOrder.Tax, updatedOrder.Tax)
	p.Suite.Equal(createOrder.Discount, updatedOrder.Discount)
	p.Suite.Equal(createOrder.TotalPrice, updatedOrder.TotalPrice)

	// ----------------------------------------------------------------------------------------------------------

	// ----------------------------------------------------------------------------------------------------------
	//! get order

	filter := make(map[string]int64)
	filter["id"] = int64(updatedOrder.Id)
	GetOrder, err := p.Repository.GetOrder(context.Background(), filter)
	p.Suite.NoError(err)
	p.Suite.NotNil(GetOrder)
	p.Suite.Equal(updatedOrder.WorkerId, GetOrder.WorkerId)
	p.Suite.Equal(updatedOrder.Tax, GetOrder.Tax)
	p.Suite.Equal(updatedOrder.Discount, GetOrder.Discount)
	p.Suite.Equal(updatedOrder.TotalPrice, GetOrder.TotalPrice)

	// ----------------------------------------------------------------------------------------------------------

	// ----------------------------------------------------------------------------------------------------------

	//! get all order

	allComment, err := p.Repository.GetOrders(context.Background(), 1, 5, map[string]string{})
	p.Suite.NoError(err)
	p.Suite.NotNil(allComment)
	// ----------------------------------------------------------------------------------------------------------

}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositrySuiteTest))
}
