package usecase

import "project/internal/entity"

type ListProdutsOutputDTO struct {
	ID    string
	Name  string
	Price float64
}

type ListProductsUsecase struct {
	ProductRepository entity.ProductRepository
}

func NewListProductsUsecase(productRepository entity.ProductRepository) *ListProductsUsecase {
	return &ListProductsUsecase{ProductRepository: productRepository}
}

func (u *ListProductsUsecase) Execute() ([]*ListProdutsOutputDTO, error) {
	p, err := u.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var productsOutput []*ListProdutsOutputDTO
	for _, product := range p {
		productsOutput = append(productsOutput, &ListProdutsOutputDTO{
			ID:    product.Name,
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return productsOutput, nil
}
