package labeler

import (
	"analyzer/drawer"
	"fmt"
)

func Run(drawer *drawer.Drawer) error {
	fmt.Println("Enter number of rows to be analyzed")
	var n uint
	_, err := fmt.Scanln(&n)
	if err != nil {
		return err
	}
	return GatherData(n, drawer)
}
