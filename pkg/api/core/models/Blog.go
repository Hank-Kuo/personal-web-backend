package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Blog struct {
	ID         int       `gorm:"primaryKey"`
	Title      string    `gorm:"NOT NULL"`
	CreateTime time.Time `gorm:"NOT NULL"`
	Tag        string    `gorm:"NOT NULL"`
	Link       string    `gorm:"NOT NULL"`
	Visitor    int       `gorm:"NOT NULL"`
	ImgLink    string    `gorm:"NOT NULL"`
	// Comments   []Comments `gorm:"ForeignKey:BlogID"`
	// Emoji      Emoji      `gorm:"ForeignKey:BlogID"`
}

func (Blog) TableName() string {
	return "Blog"
}

func BlogTable() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("Blog")
	}
}
