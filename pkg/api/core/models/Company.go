package models

import "github.com/jinzhu/gorm"

type Peo struct {
	ID        int       `gorm:"TYPE:int(11);NOT NULL;PRIMARY_KEY"`
	Name      string    `gorm:"TYPE: VARCHAR(255); DEFAULT:''"`
	CompanyID int       `gorm:"NOT NULL" sql:"type:integer REFERENCES companies(id)"`
	Company   []Company `gorm:"ForeignKey:CompanyID"`
}

type Company struct {
	ID   int    `gorm:"PRIMARY_KEY;TYPE:int(11);NOT NULL"`
	Name string `gorm:"TYPE:VARCHAR(255);DEFAULT:'';INDEX"`
	Job  string `gorm:"TYPE:VARCHAR(255);DEFAULT:''"`
}

func PeoTable() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("peos")
	}
}

func CompanyTable() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("companies")
	}
}
