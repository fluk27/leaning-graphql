package repository

import (
	"leaning-graphql/models"
	"log/slog"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// CreateProduct implements ProductRepository.
func (p productRepository) CreateProduct(product models.Product)  error{
	err:=p.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(&product)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		slog.Error("error create product", err.Error())
		return  err
	}
	return nil
}

// GetProducts implements ProductRepository.
func (p productRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	result := p.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{db: db}
}
