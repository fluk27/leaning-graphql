package services

import "leaning-graphql/models"

type ProductService interface {
	GetProductsService() ([]models.Product, error)
}
