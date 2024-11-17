package repository

import (
	"leaning-graphql/models"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// GetProducts implements ProductRepository.
func (p productRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	result:=p.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{}
}
