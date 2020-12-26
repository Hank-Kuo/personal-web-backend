package models

import "github.com/jinzhu/gorm"

type User struct {
	ID        int    `gorm:"primaryKey"`
	Account   string `gorm:"NOT NULL"`
	Password  string `gorm:"NOT NULL"`
	FirstName string `gorm:"NOT NULL"`
	LastName  string `gorm:"NOT NULL"`
	Address   string `gorm:"NOT NULL"`
	Role      string `gorm:"NOT NULL"`
}

func (User) TableName() string {
	return "User"
}

func UserTable() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("User")
	}
}
