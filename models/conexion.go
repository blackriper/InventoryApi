package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectionDatabase() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic("Error conected database" + err.Error())
	}

	db.AutoMigrate(&Product{})

	DB = db

}
