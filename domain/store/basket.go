package domain

type Basket struct {
	UserID    string  `json:"user_id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
}
