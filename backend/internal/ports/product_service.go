package ports

import (
	"wholesale-pos/internal/domain"
)

type ProductService interface {
	CreateProduct(product domain.Product) error
}
