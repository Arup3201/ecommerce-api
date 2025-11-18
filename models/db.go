package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(dataSource string) {
	var err error

	db, err = gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
