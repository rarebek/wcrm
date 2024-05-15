package entity

import (
	"time"
)

type Category struct {
	Id        string
	OwnerId   string
	Name      string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AllCategory struct {
	Categories []Category
	Count      int
}
