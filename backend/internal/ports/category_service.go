package ports

import "wholesale-pos/internal/domain"

type CategoryService interface {
	CreateCategory(category domain.Category) error
}
