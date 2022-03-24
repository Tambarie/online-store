package postgresDB

import domain "github.com/Tambarie/online-store/domain/store"

type StoreRepository interface {
	CreateProduct(product *domain.Products) (*domain.Products, error)
	GetProductDetails(reference string) (*domain.Products, error)
	UpdateProduct(reference string, products *domain.Products) error
	DeleteProduct(reference string, products *domain.Products) error
	FilterProductsByProperties(price float64, name, status string) ([]domain.Products, error)
	ListAllProducts() ([]domain.Products, error)
	AddProductToBasket(basket *domain.Basket) (*domain.Basket, error)
	GetSummaryOfProducts(basketID string) ([]domain.Basket, error)
}
