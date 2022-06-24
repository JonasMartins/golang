package main

import (
	"fmt"
	"sync"
)

func print(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"theta",
	}

	wg.Add(6)

	for i, x := range words {
		go print(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	print("Second one", &wg)
}
