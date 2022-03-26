package main

import (
	"composition/store"
	"fmt"
)

func main() {
	fmt.Println("Composition examples")

	tv := store.NewEletronic("TV", 1000, 110, false, 10)
	soundSystem := store.NewEletronic("Sound System", 1000, 220, true, 5)

	for _, i := range []*store.Eletronic{tv, soundSystem} {
		fmt.Println(i.ToString())
	}

}
