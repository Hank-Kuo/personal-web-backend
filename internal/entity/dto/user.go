package dto

import (
	"time"

	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
)

type LoginReqDto struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDto struct {
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type LoginResDto struct {
	User         model.User    `json:"user"`
	AccessToken  string        `json:"access_token"`
	Refreshtoken string        `json:"refresh_token"`
	ExpiredIn    time.Duration `json:"expired_in"`
	TokenType    string        `json:"token_type"`
}

type RegisterReqDto struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterResDto struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}
