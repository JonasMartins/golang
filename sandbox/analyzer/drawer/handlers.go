package drawer

import "fmt"

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
