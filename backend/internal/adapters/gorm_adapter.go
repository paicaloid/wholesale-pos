package adapters

import (
	"wholesale-pos/internal/domain"
	"wholesale-pos/internal/ports"

	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) ports.ProductRepository {
	return &GormProductRepository{db: db}
}

func (r *GormProductRepository) SaveProduct(product domain.Product) error {
	if result := r.db.Create(&product); result.Error != nil {
		return result.Error
	}
	return nil
}
