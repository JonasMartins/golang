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
	draws := []uint8{}
	var ball uint8
	var x int
	var alreadyWithdrown bool
	app := drawer.Drawer{
		Pot: &pot,
	}
	app.Pot = app.GeneratePot()

	for i := 0; i < 100; i++ {
		for j := 0; j < 6; j++ {
			rand.Seed(time.Now().UnixNano())
			x = rand.Intn(max-min+1) + min
			ball = uint8(x)
			alreadyWithdrown = app.CheckBallBelongs(ball, &draws)
			if alreadyWithdrown {
				for {
					x = rand.Intn(max-min+1) + min
					ball = uint8(x)
					alreadyWithdrown = app.CheckBallBelongs(ball, &draws)
					if !alreadyWithdrown {
						break
					}
				}
			}
			draws = append(draws, ball)
			app.Pot, _ = app.Draw(ball, j)
			if j == 5 {
				fmt.Printf("%02d", x)
			} else {
				fmt.Printf("%02d-", x)
			}
		}
		fmt.Print("\n")
		draws = draws[:0]
		app.Pot = app.GeneratePot()
	}
}
