package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	domain "github.com/Tambarie/online-store/domain/store"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) AddProductToBasket() gin.HandlerFunc {
	return func(context *gin.Context) {
		productID := context.Query("product_id")
		basket := &domain.Basket{}
		basket.ProductID = productID

		if err := helpers.Decode(context, basket); err != nil {
			log.Fatalf("Error  cannot bind %v", err)
			return
		}

	}
}
