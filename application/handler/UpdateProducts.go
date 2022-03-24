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

		productID := context.Param("product_id")
		product := &domain.Products{}
		product.UpdatedAt = time.Now()

		if err := helpers.Decode(context, &product); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		err := h.StoreService.UpdateProducts(productID, product)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}

		response.JSON(context, http.StatusCreated, nil, nil, "product updated successfully")
	}
}
