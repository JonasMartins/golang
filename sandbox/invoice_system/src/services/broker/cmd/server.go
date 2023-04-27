package main

import (
	"fmt"
	"time"
)

func RunLoop() {
	for {
		fmt.Println("Server Running for ever")
		time.Sleep(1 * time.Minute)
	}
}
