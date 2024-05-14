package entity

import "time"

type Owner struct {
	Id           string
	FullName     string
	CompanyName  string
	Email        string
	Password     string
	Avatar       string
	Tax          int64
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AllOwners struct {
    Owners []Owner
    Count    int
}


type CheckResponse struct {
	Check bool
}
