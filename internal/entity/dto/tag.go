package dto

import (
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type FetchTagDto struct {
	Tags []model.Tag      `json:"tags"`
	Meta utils.Pagination `json:"meta"`
}

type GetTagByIDDto struct {
	model.Tag
}
