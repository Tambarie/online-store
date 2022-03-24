package handler

import (
	"github.com/Ad3bay0c/template/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) FilterProducts() gin.HandlerFunc {
	return func(context *gin.Context) {
		price := context.Query("price")
		name := context.Query("name")
		status := context.Query("status")

		p, _ := strconv.Atoi(price)

		allProducts, err := h.StoreService.FilterProductsByProperties(float64(p), name, status)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}

		response.JSON(context, http.StatusCreated, allProducts, nil, "all products")
	}
}
