package entity

import (
	"time"
)

type Order struct {
	Id          string
	TableNumber int64
	WorkerId    string
	WorkerName  string
	Products    []ProductCheck
	Tax         int64
	Discount    int64
	TotalPrice  int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GetAllOrdersResponse struct {
	Orders     []Order
	WorkerName string
}

type ProductCheck struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Price int64  `json:"price"`
	Count int64  `json:"count"`
}

type DeleteStatus struct {
	Status bool
}
