package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	currDir, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(strings.Split(currDir, "/"))
}
