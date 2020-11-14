package dto

type LoginDto struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

type LoginReqDto struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}
