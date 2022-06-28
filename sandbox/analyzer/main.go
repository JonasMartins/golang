package main

import (
	"fmt"
	"math/rand"
	"time"

	drawer "analyzer/drawer"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	pot := make([]uint8, 60)
	app := drawer.Drawer{
		Pot: &pot,
	}
	fmt.Println("Starting Analyzer...")
	app.GenerateData(100)
}
