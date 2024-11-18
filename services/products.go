package services

import (
	"leaning-graphql/models"
	"leaning-graphql/repository"
	"log/slog"

	"github.com/google/uuid"
)

type productsServices struct {
	productRepo repository.ProductRepository
}

// CreateProductService implements ProductService.
func (p productsServices) CreateProductService(product models.Product) (models.Product, error) {
	ProductRepo:=models.Product{
		ID:          uuid.New().String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
	}
	err:=p.productRepo.CreateProduct(ProductRepo)
	if err != nil {
		return models.Product{}, err
	}
	return ProductRepo, nil
}

// GetProductsService implements ProductService.
func (p productsServices) GetProductsService() ([]models.Product, error) {
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
