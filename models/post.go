package models

import "time"

type Post struct {
	Id        uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
