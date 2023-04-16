package main

import (
	"log"
	"os"
	c "project/src/consumer"
	p "project/src/producer"
	"strings"
)

func main() {
	lenArgs, args := len(os.Args), os.Args

	if lenArgs < 2 {
		log.Fatal("need two arguments to run this program")
		os.Exit(1)
	}

	modality := strings.ToLower(args[2])

	switch modality {
	case "producer":
		p.Run()
	case "consumer":
		c.Run()
	default:
		log.Fatalf("invalid argument %s, usage: consumer | producer", modality)
		os.Exit(1)
	}
}
