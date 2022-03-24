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

func (d *RepositoryDB) GetProductDetails(reference string) (*domain.Products, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Getting products by reference : %v ...", reference))
	var product domain.Products
	err := d.db.Where("product_id = ?", reference).Find(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (d *RepositoryDB) UpdateProduct(productID string, products *domain.Products) error {
	helpers.LogEvent("INFO", fmt.Sprintf("Updating products by reference : %v ...", productID))
	result :=
		d.db.Model(domain.Products{}).
			Where("product_id = ?", productID).
			Updates(
				domain.Products{
					Model:       domain.Model{},
					ProductID:   products.ID,
					Name:        products.Name,
					Description: products.Description,
					Price:       products.Price,
					Quantity:    products.Quantity,
					Status:      products.Status,
				},
			)
	return result.Error
}

func (d *RepositoryDB) DeleteProduct(productID string, products *domain.Products) error {
	helpers.LogEvent("INFO", fmt.Sprintf("Deleting product by product_id : %v ...", productID))
	result := d.db.Where("product_id = ? ", productID).Delete(&products)
	return result.Error
}

func (d *RepositoryDB) FilterProductsByProperties(price float64, name, status string) ([]domain.Products, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Filter product by product properties : %v %v %v ...", name, price, status))
	var products []domain.Products
	if err := d.db.Where("price = ?", price).Or("name = ?", name).Or("status = ?", status).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (d *RepositoryDB) ListAllProducts() ([]domain.Products, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Listing all product by products "))
	var products []domain.Products
	if err := d.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

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
