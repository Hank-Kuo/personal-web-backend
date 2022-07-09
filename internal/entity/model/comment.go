package model

import (
	"time"
)

type Comment struct {
	ID        int       `json:"id" db:"id"`
	BlogID    int       `json:"-" db:"blog_id"`
	Name      string    `json:"name" db:"name"`
	Comment   string    `json:"comment" db:"comment"`
	Character int       `json:"character" db:"character"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
