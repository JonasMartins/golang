package labeler

import (
	"analyzer/drawer"
	"fmt"
)

type QuadrantsArrangement struct {
	Q1 uint8
	Q2 uint8
	Q3 uint8
	Q4 uint8
}

func Run(drawer *drawer.Drawer) error {
	fmt.Println("Enter number of rows to be analyzed")
	var n uint
	_, err := fmt.Scanln(&n)
	if err != nil {
		return err
	}
	return GatherData(n, drawer)
}
