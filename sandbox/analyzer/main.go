package main

import (
	drawer "analyzer/drawer"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Application need a .env file.")
	}

	pot := make([]uint8, 60)
	conn := drawer.ConnectToDB()

	app := drawer.Drawer{
		Pot: &pot,
		DB:  conn,
	}
	fmt.Println("Starting Analyzer...")
	app.GenerateData(2)
}
