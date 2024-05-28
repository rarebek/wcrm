package entity

import (
	"time"
)

type Order struct {
	Id         string
	WorkerId   string
	ProductIds []string
	Tax        int64
	Discount   int64
	TotalPrice int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DeleteStatus struct {
	Status bool
}
