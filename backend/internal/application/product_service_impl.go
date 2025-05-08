package application

import (
	"wholesale-pos/internal/domain"
	"wholesale-pos/internal/ports"
)

type ProductServiceImpl struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) CreateProduct(product domain.Product) error {
	if err := s.repo.SaveProduct(product); err != nil {
		return err
	}

	return nil
}
