package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dns := os.Getenv("DNS")
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println("Error when connecting the database: %v", err)
		return nil
	}

	return db
}
