package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	domain "github.com/Tambarie/online-store/domain/store"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) AddProductToBasket() gin.HandlerFunc {
	return func(context *gin.Context) {
		productID := context.Query("product_id")

		basket := &domain.Basket{}
		basket.ProductID = productID

		if err := helpers.Decode(context, &basket); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		_, err := h.StoreService.AddProductToBasket(basket)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}
		response.JSON(context, http.StatusCreated, nil, nil, "product added successfully")
	}
}
