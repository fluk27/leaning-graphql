package services

import "leaning-graphql/models"

type ProductService interface {
	GetProductsService() ([]models.Product, error)
	CreateProductService(product models.Product) (models.Product, error)
}
