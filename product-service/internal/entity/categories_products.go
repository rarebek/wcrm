package entity

import (
	"time"
)

type CategoryProduct struct {
	Id         int64
	ProductId  int64
	CategoryId int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

