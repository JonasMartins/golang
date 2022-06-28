package cmd

import (
	drawer "analyzer/drawer"
	"errors"
	"fmt"
	"math"
)

func GenerateNewData() error {

	pot := make([]uint8, 60)
	conn := drawer.ConnectToDB()

	fmt.Println("Enter a number of rows to be inserted into database")
	var n uint
	_, err := fmt.Scanln(&n)
	if err != nil {
		return err
	}
	app := drawer.Drawer{
		Pot: &pot,
		DB:  conn,
	}
	app.GenerateData(int(math.Min(10, float64(n))))
	fmt.Println("Operation done successfully")
	return errors.New("continue")
}
