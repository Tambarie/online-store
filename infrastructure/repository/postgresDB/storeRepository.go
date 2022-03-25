package postgresDB

import (
	"fmt"
	"github.com/Tambarie/online-store/application/helpers"
	domain "github.com/Tambarie/online-store/domain/store"
	"gorm.io/gorm"
	"log"
)

// RepositoryDB struct
type RepositoryDB struct {
	db *gorm.DB
}

// NewPaymentGatewayRepositoryDB function to initialize RepositoryDB struct
func NewPaymentGatewayRepositoryDB(client *gorm.DB) *RepositoryDB {
	return &RepositoryDB{
		db: client,
	}
}

// CreateProduct Saves the product to the database
func (d *RepositoryDB) CreateProduct(product *domain.Products) (*domain.Products, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Creating products : %v ...", product))
	err := d.db.Create(&product).Error
	log.Println("product created")
	if err != nil {
		return nil, err
	}
	log.Println("COMMENT HERE", product)
	return product, nil
}

// GetProductDetails Gets product by the ID
func (d *RepositoryDB) GetProductDetails(productID string) (*domain.Products, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Getting products by productID : %v ...", productID))
	var product domain.Products
	err := d.db.Where("product_id = ?", productID).Find(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct Update the product in the database
func (d *RepositoryDB) UpdateProduct(productID string, products *domain.Products) error {
	helpers.LogEvent("INFO", fmt.Sprintf("Updating products by reference : %v ...", productID))
	result :=
		d.db.Model(domain.Products{}).
			Where("product_id = ?", productID).
			Updates(
				domain.Products{
					Model:       domain.Model{},
					ProductID:   products.ProductID,
					Name:        products.Name,
					Description: products.Description,
					Price:       products.Price,
					Quantity:    products.Quantity,
					Status:      products.Status,
				},
			)
	return result.Error
}

// DeleteProduct Delete the product in the database
func (d *RepositoryDB) DeleteProduct(productID string, products *domain.Products) error {
	helpers.LogEvent("INFO", fmt.Sprintf("Deleting product by product_id : %v ...", productID))
	result := d.db.Where("product_id = ? ", productID).Delete(&products)
	return result.Error
}

// FilterProductsByProperties filters the product by it's properties
func (d *RepositoryDB) FilterProductsByProperties(price float64, name, status string) ([]domain.Products, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Filter product by product properties : %v %v %v ...", name, price, status))
	var products []domain.Products
	if err := d.db.Where("price = ?", price).Or("name = ?", name).Or("status = ?", status).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// ListAllProducts List all the products
func (d *RepositoryDB) ListAllProducts() ([]domain.Products, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Listing all product by products "))
	var products []domain.Products
	if err := d.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// AddProductToBasket Adds product to the basket
func (d *RepositoryDB) AddProductToBasket(basket *domain.Basket) (*domain.Basket, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Adding  product to basket %v ...", basket))
	err := d.db.Create(&basket).Error
	log.Println("product created")
	if err != nil {
		return nil, err
	}
	log.Println("COMMENT HERE", basket)
	return basket, nil
}

// GetSummaryOfProducts Gets the summary of the product
func (d *RepositoryDB) GetSummaryOfProducts(basketID string) ([]domain.Basket, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Getting the summary of products in basket with basketId %v ", basketID))
	var basket []domain.Basket
	if err := d.db.Where("user_id = ?", basketID).Find(&basket).Error; err != nil {
		return nil, err
	}
	return basket, nil
}
