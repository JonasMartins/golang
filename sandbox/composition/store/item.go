package store

import "fmt"

type Item struct {
	Name, Category string
	price          float64
}

func NewItem(name, category string, price float64) *Item {
	return &Item{name, category, price}
}

func (i *Item) Price() float64 {
	return i.price
}

func (i *Item) ToString() string {
	return fmt.Sprintf("Name: %s, Category: %s, Price %.2f", i.Name, i.Category, i.Price())
}
