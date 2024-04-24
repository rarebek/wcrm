package entity

import (
	"time"
)

type Order struct {
	Id         int64
	WorkerId   string
	ProductId  int64
	Tax        int64
	Discount   int64
	TotalPrice int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
