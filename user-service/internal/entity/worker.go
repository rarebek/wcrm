package entity

import "time"

type Worker struct {
	Id        string
	FullName  string
	LoginKey  string
	Password  string
	OwnerId   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
