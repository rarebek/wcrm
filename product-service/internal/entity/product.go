package entity

import (
	"time"

)

type Product struct {
	Id          int64
	Title       string
	Description string
	Price       int64
	Discount    int64
	Picture     string
	CategoryId  int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type AllProduct struct {
    Products []Product
    Count    int
}

type CheckResponse struct {
	Check bool
}
