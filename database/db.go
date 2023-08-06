package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/models"
)

var DB *gorm.DB

func ConnectToDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/task5_vix_btpns"))

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Photo{})

	if err != nil {
		panic(err)
	}

	DB = db
}