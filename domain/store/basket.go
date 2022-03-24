package domain

type Basket struct {
	Model
	ProductID string  `json:"id"`
	Price     float64 `json:"price"`
}
