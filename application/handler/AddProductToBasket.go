package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	domain "github.com/Tambarie/online-store/domain/store"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// AddProductToBasket Handles the addition of products to the basket
func (h *Handler) AddProductToBasket() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Logging the context
		helpers.LogRequest(context)

		// Getting the productID from the query string
		productID := context.Query("product_id")

		basket := &domain.Basket{}
		basket.ProductID = productID

		// Binding the json
		if err := helpers.Decode(context, &basket); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		// saving products to basket
		_, err := h.StoreService.AddProductToBasket(basket)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}

		// JSON Response to the client
		response.JSON(context, http.StatusCreated, nil, nil, "product added successfully")
	}
}
