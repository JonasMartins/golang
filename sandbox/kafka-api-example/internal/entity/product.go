package entity

type Product struct {
	ID    uint32
	Name  string
	Price float64
}

type ProductRepository interface {
	Create(product *Product) error
	FindAll() ([]*Product, error)
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    0,
		Name:  name,
		Price: price,
	}
}
