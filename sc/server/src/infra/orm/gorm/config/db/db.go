package db

import (
	"fmt"
	"log"
	"os"

	"src/infra/orm/gorm/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Application need a .env file.")
	}
	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_pass := os.Getenv("DB_PASS")

	// source: https://dev.to/karanpratapsingh/connecting-to-postgresql-using-gorm-24fj
	// https://gorm.io/docs/

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_pass, db_name, "5432")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Warn),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database successfully connected")
	}

	db.AutoMigrate(&models.User{}, &models.Chat{}, &models.Message{})

	return db
}
