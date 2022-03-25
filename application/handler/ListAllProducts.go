package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ListAllProducts() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Logging the context
		helpers.LogRequest(context)

		// List all products in the database
		allProducts, err := h.StoreService.ListAllProducts()
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}
		// JSON Response to the client
		response.JSON(context, http.StatusCreated, allProducts, nil, "all products")
	}
}
