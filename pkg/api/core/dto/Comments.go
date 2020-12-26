package dto

import (
	"time"
)

type GetCommentsResDto struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Comment    string    `json:"comment"`
	CreateTime time.Time `json:"create_time"`
	Character  int       `json:"character"`
}

type PostCommentsReqDto struct {
	BlogID     int       `json:"blog_id" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Comment    string    `json:"comment" binding:"required"`
	CreateTime time.Time `json:"create_time"`
	Character  int       `json:"character"`
}

type PostCommentsResDto struct {
}
