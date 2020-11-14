package dto

type PeopleDto struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	City      string `json:"city" binding:"required"`
}

type GetAllPeopleReqDto struct {
}

type GetAllPeopleResDto struct {
}

type GetPeopleReqDto struct {
}

type GetPeopleResDto struct {
}
