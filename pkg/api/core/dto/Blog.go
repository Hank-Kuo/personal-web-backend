package dto

import "time"

type GetAllBlogResDto struct {
	ID         uint      `json: "id"`
	Title      string    `json: "title"`
	CreateTime time.Time `json: "create_time"`
	Tag        string    `json: "tag"`
	Visitor    int       `json: "visitor"`
	Emoji      int       `json: "emoji"`
	Like       int       `json: "like"`
	ImgLink    string    `json: "img_link"`
}

type GetBlogResDto struct {
	ID   uint   `json:"id"`
	Link string `json:"link"`
}
