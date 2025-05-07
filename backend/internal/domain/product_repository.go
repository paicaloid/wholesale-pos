package domain

type ProductRepository interface {
	SaveProduct(product Product) error
}
