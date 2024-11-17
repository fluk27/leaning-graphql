package repository

import (
	"leaning-graphql/models"
)

type ProductRepository interface {
	GetProducts() ([]models.Product, error)
}
