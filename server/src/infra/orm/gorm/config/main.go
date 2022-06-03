package config

import (
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

import Init() *gorm.DB {

	_ := godotenv.Load()

	db_user := os.Getenv("DB_USER")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")
	db_pass := os.Getenv("DB_PASS")

	// source: https://dev.to/karanpratapsingh/connecting-to-postgresql-using-gorm-24fj
	// https://gorm.io/docs/

	url := fmt.Printf()
}