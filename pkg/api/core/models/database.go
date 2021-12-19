package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func ConnectDB(adapter string, host string) *gorm.DB {
	db, err = gorm.Open(adapter, host)
	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)
	db.AutoMigrate(&User{}, &Comments{}, &Emoji{}, &Blog{})
	return db
}

func CloseDB() error {
	return db.Close()
}
