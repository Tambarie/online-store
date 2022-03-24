package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetSummaryInBasket() gin.HandlerFunc {
	return func(context *gin.Context) {
		var totalPrice float64
		user := struct {
			UserID string `json:"user_id"`
		}{}

		if err := helpers.Decode(context, &user); err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}

		basketDB, err := h.StoreService.GetSummaryOfProducts(user.UserID)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}

		for _, basket := range basketDB {
			totalPrice += basket.Price
		}
		response.JSON(context, http.StatusOK, gin.H{
			"count": len(basketDB),
			"price": totalPrice,
		}, nil, "")
	}
}
