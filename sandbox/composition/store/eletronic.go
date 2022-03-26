package store

type Eletronic struct {
	*Item
	voltage   float64
	battery   bool
	lifeLimit int
}

func NewEletronic(name string, price float64, voltage float64, battery bool, lifeLimit int) *Eletronic {
	return &Eletronic{
		NewItem(name, "Eletronic", price), voltage, battery, lifeLimit,
	}
}
