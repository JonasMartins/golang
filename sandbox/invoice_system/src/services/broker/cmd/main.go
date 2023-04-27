package main

import (
	"fmt"
	cfg "project/src/services/broker/configs"
)

func main() {
	fmt.Println("Hello from broker")

	cfg, err := cfg.LoadConfig()
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("cfg: %v", cfg)
	}

	RunLoop()

}
