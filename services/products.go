package services

import (
	"leaning-graphql/models"
	"leaning-graphql/repository"
	"log/slog"
)

type productsServices struct {
	productRepo repository.ProductRepository
}

// GetProductsService implements ProductService.
func (p productsServices) GetProductsService() ([]models.Product, error) {
	return []models.Product{
		{
			ID:          "1",
			Name:        "Product 1",
			Description: "Description 1",
			Price:       100,
			Image:       nil,
		},
		{
			ID:          "2",
			Name:        "Product 2",
			Description: "Description 2",
			Price:       200,
			Image:       nil,
		},
	}, nil
	products, err := p.productRepo.GetProducts()
	if err != nil {
		slog.Error("error get products", err.Error())
		return nil, err
	}
	return products, nil
}

func NewProductServices(productRepo repository.ProductRepository) ProductService {
	return productsServices{
		productRepo: productRepo,
	}
}
