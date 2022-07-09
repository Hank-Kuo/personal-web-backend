package model

import "time"

type User struct {
	ID         int       `json:"id" db:"id"`
	UUID       string    `json:"uuid" db:"uuid"`
	Account    string    `json:"account" db:"account"`
	Password   string    `json:"-" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Email      string    `json:"email" db:"email"`
	Role       string    `json:"role" db:"role"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Login_time time.Time `json:"login_time" db:"login_time"`
}
