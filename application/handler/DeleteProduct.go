package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	domain "github.com/Tambarie/online-store/domain/store"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) DeleteProduct() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Logging the context
		helpers.LogRequest(context)

		product := struct {
			ProductID string `json:"product_id"`
		}{}

		products := &domain.Products{}

		// Binds the JSON
		if err := helpers.Decode(context, &product); err != nil {
			log.Fatalf("Error  cannot bind %v", err)
			return
		}

		// Deletes product from the database
		err := h.StoreService.DeleteProduct(product.ProductID, products)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}

		// JSON Response to the client
		response.JSON(context, http.StatusOK, nil, nil, "product delete successfully")

	}
}
