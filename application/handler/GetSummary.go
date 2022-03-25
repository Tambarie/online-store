package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetSummaryInBasket() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Logging the context
		helpers.LogRequest(context)

		var totalPrice float64
		user := struct {
			UserID string `json:"user_id"`
		}{}

		// Binding the json
		if err := helpers.Decode(context, &user); err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}

		// Get's the summary of the products in a basket
		basketDB, err := h.StoreService.GetSummaryOfProducts(user.UserID)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}

		// sums the total
		for _, basket := range basketDB {
			totalPrice += basket.Price
		}
		response.JSON(context, http.StatusOK, gin.H{
			"Total count":     len(basketDB),
			"Total price":     totalPrice,
			"created_baskets": basketDB,
		}, nil, "")
	}
}
