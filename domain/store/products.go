package domain

type ProductStatus string

const (
	INSTOCK     ProductStatus = "INSTOCK"
	ORDERED     ProductStatus = "ORDERED"
	UNAVAILABLE ProductStatus = "UNAVAILABLE"
)

type Products struct {
	Model
	ProductID   string        `json:"product_id"`
	Name        string        `json:"name" `
	Description string        `json:"description"`
	Price       float64       `json:"price"`
	Quantity    int           `json:"quantity"`
	Status      ProductStatus `json:"status"`
}
