package dto

type LoginReqDto struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResDto struct {
	Account   string `json:"account"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

type RegisterReqDto struct {
	Account   string `json:"account" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Role      string `json:"role"`
}

type SerialLoginResDto struct {
	Account   string `json:"account"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

func SerializeLoginDto(data *LoginResDto) *SerialLoginResDto {
	return &SerialLoginResDto{
		Account:   data.Account,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Role:      data.Role,
		Token:     data.Token,
	}
}
