package database

import (
	"fmt"
	"os"
	"github.com/devanfer02/task-5-vix-btpns-devanferrel/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	err := godotenv.Load();

	if err != nil {
		panic(err)
	}	

	dsn := fmt.Sprintf(
		"%s%s:@tcp(%s)/%s", 
		os.Getenv("DB_USERNAME"), 
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_CONNECTION"), 
		os.Getenv("DB_NAME"),
	)
	
	db, err := gorm.Open(mysql.Open(dsn))

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