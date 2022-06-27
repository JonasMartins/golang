package main

import (
	"fmt"
	"math/rand"
	"time"

	drawer "analyzer/drawer"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	min := 1
	max := drawer.MAX

	pot := make([]uint8, 60)
	draws := make([]uint8, 6)
	var ball uint8
	var x int
	var alreadyWithdrown bool
	app := drawer.Drawer{
		Pot: &pot,
	}
	app.Pot = app.GeneratePot()
	fmt.Println("Length ", len(*app.Pot))
	x = 10

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			x = rand.Intn(max-min+1) + min
			ball = uint8(x)
			alreadyWithdrown = app.CheckBallBelongs(ball, &draws)
			if alreadyWithdrown {
				for !alreadyWithdrown {
					x = rand.Intn(max-min+1) + min
					ball = uint8(x)
					alreadyWithdrown = app.CheckBallBelongs(ball, &draws)
				}
			}
			draws = append(draws, ball)
			app.Pot, _ = app.Draw(ball, j)
			fmt.Printf("%02d-", x)
		}
		x = rand.Intn(max-min+1) + min
		ball = uint8(x)
		alreadyWithdrown = app.CheckBallBelongs(ball, &draws)
		if alreadyWithdrown {
			for !alreadyWithdrown {
				x = rand.Intn(max-min+1) + min
				ball = uint8(x)
				alreadyWithdrown = app.CheckBallBelongs(ball, &draws)
			}
		}
		draws = append(draws, ball)
		app.Pot, _ = app.Draw(ball, 5)
		fmt.Printf("%02d", x)
		fmt.Print("\n")
		fmt.Println("Length ", len(*app.Pot))
		app.Pot = app.GeneratePot()
	}

}
