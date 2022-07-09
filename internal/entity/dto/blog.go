package dto

import (
	"github.com/Hank-Kuo/personal-web-backend/internal/entity/model"
	"github.com/Hank-Kuo/personal-web-backend/pkg/utils"
)

type BlogDto struct {
	model.Blog
	Tags []model.Tag `json:"tags"`
}

type FetchBlogDto struct {
	Blogs []BlogDto        `json:"blogs"`
	Meta  utils.Pagination `json:"meta"`
}

type GetBlogByIDDto struct {
	BlogDto
	Emoji model.Emoji `json:"emoji"`
	Html  string      `json:"html"`
}

type CreateBlogDto struct {
	Title   string   `json:"title" binding:"required"`
	Link    string   `json:"link" binding:"required"`
	ImgLink string   `json:"img_link" binding:"required"`
	Tags    []string `json:"tags" binding:"required"`
}

type UpdateBlogDto struct {
	Title   string   `json:"title,omitempty"`
	Link    string   `json:"link,omitempty"`
	ImgLink string   `json:"img_link,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Visitor int      `json:"visitor,omitempty"`
}
