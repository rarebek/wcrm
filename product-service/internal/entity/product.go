package entity

import (
	"time"
)

type Product struct {
	Id          string
	OwnerId     string
	Title       string
	Description string
	Price       int64
	Discount    int64
	Picture     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type ProductWithCategoryId struct {
	Id          string
	OwnerId     string
	Title       string
	Description string
	Price       int64
	Discount    int64
	Picture     string
	CategoryId  string
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
