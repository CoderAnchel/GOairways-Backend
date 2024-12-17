package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func StartDB() {
	dsn := "root:kgyspy10230@tcp(localhost:3306)/aviones?parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection successfully opened!")
	DB = db
}

func CloseDB() {
	db, err := DB.DB()

	if err != nil {
		panic("Failed to close database!")
	}

	db.Close()
}
