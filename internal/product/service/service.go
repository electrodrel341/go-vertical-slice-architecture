package service

import (
	"PetAi/internal/product/infrastructure"
	"go.uber.org/dig"
)

// product service instance
type ProductService struct {
	CreateProductServiceProvider *CreateProductService
}

func NewProductService() *ProductService {
	return &ProductService{}
}

// provide components for injection
func ProvideProductComponents(c *dig.Container) {
	// repositorory provider injection
	err := c.Provide(infrastructure.NewProductRepository)
	if err != nil {
		panic(err)
	}

	// service provider injection
	err = c.Provide(NewCreateProductService)
	if err != nil {
		panic(err)
	}
}

// init service container
func (ps *ProductService) InitProductComponents(c *dig.Container) error {
	// create product service
	err := c.Invoke(
		func(s *CreateProductService) {
			ps.CreateProductServiceProvider = s
		},
	)

	return err
}
