package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Blog struct {
	ID         uint      `json: "id"`
	Title      string    `json: "title"`
	CreateTime time.Time `json: "create_time"`
	Tag        string    `json: "tag"`
	Link       string    `json: "link"`
	Visitor    int       `json: "visitor"`
	Emoji      int       `json: "emoji"`
	Like       int       `json: "like"`
	ImgLink    string    `json: "img_link"`
}

func (Blog) TableName() string {
	return "Blog"
}

func BlogTable() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("Blog")
	}
}
