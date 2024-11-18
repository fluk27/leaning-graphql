package repository

import (
	"leaning-graphql/models"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
	CreateProduct(product models.Product) error
}
