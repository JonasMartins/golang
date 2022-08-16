package main

import (
	"os"
	application "src/main/config"
)

func main() {
	application.Run()
}

func GetCurrPath() string {
	rootPath, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	return rootPath
}
