package model

import "time"

type Blog struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Link      string    `json:"link" db:"link"`
	Visitor   int       `json:"visitor" db:"visitor"`
	ImgLink   string    `json:"img_link" db:"img_link"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type BlogToTag struct {
	ID   int   `json:"id" db:"id"`
	Tags []Tag `json:"json" db:"tags"`
}
