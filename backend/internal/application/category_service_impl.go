package application

import (
	"wholesale-pos/internal/domain"
	"wholesale-pos/internal/ports"
)

type CategoryServiceImpl struct {
	repo ports.CategoryRepository
}

func NewCategoryService(repo ports.CategoryRepository) ports.CategoryService {
	return &CategoryServiceImpl{repo: repo}
}

func (s *CategoryServiceImpl) CreateCategory(category domain.Category) error {
	if err := s.repo.SaveCatogory(category); err != nil {
		return err
	}

	return nil
}
