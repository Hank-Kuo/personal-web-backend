package models

import "github.com/jinzhu/gorm"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement:false"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Address   string `json:"address"`
	Role      string `json:"role"`
}

func UserTable() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("users")
	}
}
