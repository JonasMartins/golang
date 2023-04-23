package usecase

import "project/internal/entity"

type CreateProductInputDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateProductOutputDTO struct {
	ID    uint32
	Name  string
	Price float64
}

type CreateProductUsecase struct {
	ProductRepository entity.ProductRepository
}

func NewCreateProductUsecase(productRepository entity.ProductRepository) *CreateProductUsecase {
	return &CreateProductUsecase{ProductRepository: productRepository}
}

func (u *CreateProductUsecase) Execute(
	input *CreateProductInputDTO,
) (*CreateProductOutputDTO, error) {
	p := entity.NewProduct(input.Name, input.Price)
	err := u.ProductRepository.Create(p)
	if err != nil {
		return nil, err
	}
	return &CreateProductOutputDTO{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}, nil
}
