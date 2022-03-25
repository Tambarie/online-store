package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	domain "github.com/Tambarie/online-store/domain/store"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) UpdateProducts() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Logging the context
		helpers.LogRequest(context)

		productID := context.Param("product_id")
		product := &domain.Products{}
		product.UpdatedAt = time.Now()

		// Binds the JSON
		if err := helpers.Decode(context, &product); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		// Updates the product by it's ID
		err := h.StoreService.UpdateProducts(productID, product)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}

		// JSON Response to the client
		response.JSON(context, http.StatusCreated, nil, nil, "product updated successfully")
	}
}
