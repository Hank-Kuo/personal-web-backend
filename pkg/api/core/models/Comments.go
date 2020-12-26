package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comments struct {
	ID         int       `gorm:"primaryKey"`
	BlogID     int       `gorm:"NOT NULL" sql:"type:integer REFERENCES Blog(id)"`
	Name       string    `gorm:"NOT NULL"`
	Comment    string    `gorm:"NOT NULL"`
	CreateTime time.Time `gorm:"NOT NULL"`
	Character  int       `gorm:"NOT NULL"`
}

func (Comments) TableName() string {
	return "Comments"
}

func CommentsTable() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("Comments")
	}
}
