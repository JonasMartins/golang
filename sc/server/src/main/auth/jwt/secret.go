package jwt

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetJwtSecret() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Application need a .env file.")
	}

	secret := os.Getenv("PORT")

	return []byte(secret)
}
