package service

import (
	domain "github.com/Tambarie/online-store/domain/store"
	"github.com/Tambarie/online-store/infrastructure/repository/postgresDB"
)

type StoreService interface {
	CreateProduct(product *domain.Products) (*domain.Products, error)
	GetProductDetails(reference string) (*domain.Products, error)
	UpdateProducts(reference string, products *domain.Products) error
	DeleteProduct(reference string, products *domain.Products) error
	FilterProductsByProperties(price float64, name, status string) ([]domain.Products, error)
	ListAllProducts() ([]domain.Products, error)
	AddProductToBasket(basket *domain.Basket) (*domain.Basket, error)
}

type DefaultAccountService struct {
	repo postgresDB.StoreRepository
}

func NewStoreService(repo postgresDB.StoreRepository) *DefaultAccountService {
	return &DefaultAccountService{
		repo: repo,
	}
}

func (s *DefaultAccountService) CreateProduct(product *domain.Products) (*domain.Products, error) {
	return s.repo.CreateProduct(product)
}

func (s *DefaultAccountService) GetProductDetails(reference string) (*domain.Products, error) {
	return s.repo.GetProductDetails(reference)
}

func (s DefaultAccountService) UpdateProducts(reference string, products *domain.Products) error {
	return s.repo.UpdateProduct(reference, products)
}

func (s *DefaultAccountService) DeleteProduct(reference string, products *domain.Products) error {
	return s.repo.DeleteProduct(reference, products)
}

func (s *DefaultAccountService) FilterProductsByProperties(price float64, name, status string) ([]domain.Products, error) {
	return s.repo.FilterProductsByProperties(price, name, status)
}

func (s *DefaultAccountService) ListAllProducts() ([]domain.Products, error) {
	return s.repo.ListAllProducts()
}

func (s *DefaultAccountService) AddProductToBasket(basket *domain.Basket) (*domain.Basket, error)  {
	return s.repo.AddProductToBasket(basket)
}