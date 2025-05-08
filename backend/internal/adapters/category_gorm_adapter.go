package adapters

import (
	"wholesale-pos/internal/domain"
	"wholesale-pos/internal/ports"

	"gorm.io/gorm"
)

type GormCategoryRepository struct {
	db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) ports.CategoryRepository {
	return &GormCategoryRepository{db: db}
}

func (r *GormCategoryRepository) SaveCatogory(category domain.Category) error {
	if result := r.db.Create(&category); result.Error != nil {
		return result.Error
	}
	return nil
}
