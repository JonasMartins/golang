package drawer

import (
	"fmt"
	"math/rand"
	"time"
)

func (d *Drawer) GeneratePot() *[]uint8 {
	d.Pot = nil
	pot := []uint8{}
	var j uint8

	for j = 1; j <= MAX; j++ {
		pot = append(pot, j)
	}

	return &pot
}

func (d *Drawer) Draw(n uint8, round int) (*[]uint8, error) {

	if n > MAX {
		return nil, fmt.Errorf("invalid number to draw")
	}

	index := d.Find(n)

	if index >= (MAX - round) {
		return nil, fmt.Errorf("number already withdrawn")
	} else {
		d.Pot = d.WithDraw(index)
	}

	return d.Pot, nil

}

// Finds a number from pot and return its index
// If not Found, return MAX + 1
func (d *Drawer) Find(n uint8) int {
	//fmt.Printf("\npot length %d\n", len(*d.Pot))
	for i, x := range *d.Pot {
		if n == x {
			return i
		}
	}

	return MAX + 1
}

// make sure the number to withdraw exists
// then withdraw it from the pot then
// the pot gets rebuilt
func (d *Drawer) WithDraw(n int) *[]uint8 {

	var aux = *d.Pot
	aux = append(aux[:n], aux[n+1:]...)
	return &aux
}

// true if a int belongs to an array of ints
func (d *Drawer) CheckBallBelongs(n uint8, arr *[]uint8) bool {
	for _, x := range *arr {
		if n == x {
			return true
		}
	}
	return false
}

func (d *Drawer) GenerateData(amount int) {

	min := 1
	max := MAX
	draws := []uint8{}
	var ball uint8
	var x int
	var alreadyWithdrown bool

	d.Pot = d.GeneratePot()

	for i := 0; i < amount; i++ {
		for j := 0; j < 6; j++ {
			rand.Seed(time.Now().UnixNano())
			x = rand.Intn(max-min+1) + min
			ball = uint8(x)
			alreadyWithdrown = d.CheckBallBelongs(ball, &draws)
			if alreadyWithdrown {
				for {
					x = rand.Intn(max-min+1) + min
					ball = uint8(x)
					alreadyWithdrown = d.CheckBallBelongs(ball, &draws)
					if !alreadyWithdrown {
						break
					}
				}
			}
			draws = append(draws, ball)
			d.Pot, _ = d.Draw(ball, j)
			if j == 5 {
				fmt.Printf("%02d", x)
			} else {
				fmt.Printf("%02d-", x)
			}
		}
		fmt.Print("\n")
		draws = draws[:0]
		d.Pot = d.GeneratePot()
	}

}
