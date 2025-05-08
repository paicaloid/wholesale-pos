package ports

import (
	"wholesale-pos/internal/domain"
)

type ProductRepository interface {
	SaveProduct(product domain.Product) error
}
