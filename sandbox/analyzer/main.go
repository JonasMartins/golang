package main

import (
	cmd "analyzer/cmd"
	//labeler "analyzer/labeler"
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

	//d := cmd.StartDrawer()

	//labeler.GatherData(100, d)

	cmd.Run()

}
