package cmd

import (
	drawer "analyzer/drawer"
	labeler "analyzer/labeler"
	"errors"
	"fmt"
	"math"
)

func StartDrawer() *drawer.Drawer {

	pot := make([]uint8, 60)
	conn := drawer.ConnectToDB()

	app := drawer.Drawer{
		Pot: &pot,
		DB:  conn,
	}

	return &app
}

func GenerateNewData() error {

	app := StartDrawer()

	fmt.Println("Enter a number of rows to be inserted into database")
	var n uint
	_, err := fmt.Scanln(&n)
	if err != nil {
		return err
	}
	app.GenerateData(int(math.Min(100, float64(n))))
	fmt.Println("Operation done successfully")
	return errors.New("continue")
}

func StartAnalysis() error {

	app := StartDrawer()
	return labeler.Run(app)

}
