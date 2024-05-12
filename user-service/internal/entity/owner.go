package entity

import "time"

type Owner struct {
	Id          string
	FullName    string
	CompanyName string
	Email       string
	Password    string
	Avatar      string
	Tax         int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CheckResponse struct {
	Check bool
}

