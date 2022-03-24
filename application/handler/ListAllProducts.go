package handler

import (
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ListAllProducts() gin.HandlerFunc {
	return func(context *gin.Context) {

		allProducts, err := h.StoreService.ListAllProducts()
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}
		response.JSON(context, http.StatusCreated, allProducts, nil, "all products")
	}
}
