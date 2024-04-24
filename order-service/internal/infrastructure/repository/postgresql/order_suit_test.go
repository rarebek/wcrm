package postgresql

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	pb "order-service/genproto/order"
	"order-service/internal/pkg/config"
	db "order-service/internal/pkg/postgres"

	model "order-service/internal/entity"
	Order "order-service/internal/infrastructure/repository"
)

type OrderRepositrySuiteTest struct {
	suite.Suite
	Repository  Order.Order
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
	
	

	productReq := &pb.Order{

	}
	createdProduct, err := p.Repository.CreateProduct(context.Background(), &model.Product{
		Title: productReq.Title,
		Description: productReq.Description,
		Price: productReq.Price,
		Discount: productReq.Discount,
		Picture: productReq.Picture,
		CategoryId: productReq.CategoryId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	p.Suite.NoError(err)
	p.Suite.NotNil(createdProduct)
	p.Suite.NotNil(createdProduct.Id)
	p.Suite.Equal(productReq.Title, createdProduct.Title)
	p.Suite.Equal(productReq.Description, createdProduct.Description)
	p.Suite.Equal(productReq.Price, createdProduct.Price)
	p.Suite.Equal(productReq.Discount, createdProduct.Discount)
	p.Suite.Equal(productReq.Picture, createdProduct.Picture)
	p.Suite.Equal(productReq.CategoryId, createdProduct.CategoryId)
	p.Suite.NotNil(createdProduct.CreatedAt)
	p.Suite.NotNil(createdProduct.UpdatedAt)

	// ---------------------------------------------------------------------------------------------------------


	// update product
	// ---------------------------------------------------------------------------------------------------------

	createdProduct.Title = "updated title"
	createdProduct.Description = "update description"
	createdProduct.Price = 10000
	createdProduct.Discount = 10
	createdProduct.Picture = "updated/picture.pdf"
	createdProduct.CategoryId = 12345

	updatedProduct, err := p.Repository.UpdateProduct(context.Background(), createdProduct)
	p.Suite.NoError(err)
	p.Suite.NotNil(updatedProduct)
	p.Suite.NotNil(updatedProduct.Id)
	p.Suite.Equal(createdProduct.Title, updatedProduct.Title)
	p.Suite.Equal(createdProduct.Description, updatedProduct.Description)
	p.Suite.Equal(createdProduct.Price, updatedProduct.Price)
	p.Suite.Equal(createdProduct.Discount, updatedProduct.Discount)
	p.Suite.Equal(createdProduct.Picture, updatedProduct.Picture)
	p.Suite.Equal(createdProduct.CategoryId, updatedProduct.CategoryId)
	
	// ----------------------------------------------------------------------------------------------------------


	// ----------------------------------------------------------------------------------------------------------
	// get product

	filter := make(map[string]int64)
	filter["id"] = int64(updatedProduct.Id)
	GetProduct, err := p.Repository.GetProduct(context.Background(), filter)
	p.Suite.NoError(err)
	p.Suite.NotNil(GetProduct)
	p.Suite.Equal(updatedProduct.Title, GetProduct.Title)
	p.Suite.Equal(updatedProduct.Description, GetProduct.Description)
	p.Suite.Equal(updatedProduct.Price, GetProduct.Price)
	p.Suite.Equal(updatedProduct.Discount, GetProduct.Discount)
	p.Suite.Equal(updatedProduct.Picture, GetProduct.Picture)
	p.Suite.Equal(updatedProduct.CategoryId, GetProduct.CategoryId)
	// ----------------------------------------------------------------------------------------------------------


	// ----------------------------------------------------------------------------------------------------------

	// get all comment
	allComment, err := p.Repository.ListProduct(context.Background(), 1, 5, map[string]string{})
	p.Suite.NoError(err)
	p.Suite.NotNil(allComment)
	// ----------------------------------------------------------------------------------------------------------
	

	// ----------------------------------------------------------------------------------------------------------
	// delete user
	delproduct, err := p.Repository.DeleteProduct(context.Background(),createdProduct.Id) 
	p.Suite.NoError(err)
	p.Suite.NotNil(delproduct)
	// ----------------------------------------------------------------------------------------------------------
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositrySuiteTest))
}
