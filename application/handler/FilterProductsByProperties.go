package handler

import (
	"github.com/Ad3bay0c/template/response"
	"github.com/Tambarie/online-store/application/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) FilterProducts() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Logging the context
		helpers.LogRequest(context)

		// Getting the query parameters
		price := context.Query("price")
		name := context.Query("name")
		status := context.Query("status")

		//  converting the price to int
		p, err := strconv.Atoi(price)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}

		// method that filterer's a product by its properties
		allProducts, err := h.StoreService.FilterProductsByProperties(float64(p), name, status)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}

		// JSON Response to the client
		response.JSON(context, http.StatusCreated, allProducts, nil, "all products")
	}
}
