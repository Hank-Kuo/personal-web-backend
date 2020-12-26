package models

import (
	"github.com/jinzhu/gorm"
)

type Emoji struct {
	ID      int `gorm:"primaryKey"`
	BlogID  int `gorm:"NOT NULL" sql:"type:integer REFERENCES Blog(id)"`
	Funny   int `gorm:"NOT NULL"`
	Sad     int `gorm:"NOT NULL"`
	Wow     int `gorm:"NOT NULL"`
	Clap    int `gorm:"NOT NULL"`
	Perfect int `gorm:"NOT NULL"`
	Love    int `gorm:"NOT NULL"`
	Hard    int `gorm:"NOT NULL"`
	Good    int `gorm:"NOT NULL"`
	Mad     int `gorm:"NOT NULL"`
}

func (Emoji) TableName() string {
	return "Emoji"
}

func EmojiTable() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("Emoji")
	}
}
