package dto

import (
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type GetCommentByBlogIDDto struct {
	Comment []model.Comment  `json:"comments"`
	Meta    utils.Pagination `json:"meta"`
}

type CraeteCommentDto struct {
	BlogID  int    `json:"blog_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}
