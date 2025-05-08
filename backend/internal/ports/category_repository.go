package ports

import "wholesale-pos/internal/domain"

type CategoryRepository interface {
	SaveCatogory(domain.Category) error
}
