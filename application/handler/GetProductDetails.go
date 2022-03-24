package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) GetProductDetails() gin.HandlerFunc {
	return func(context *gin.Context) {
		product := struct {
			ProductID string `json:"product_id"`
		}{}

		if err := helpers.Decode(context, &product); err != nil {
			log.Fatalf("Error %v", err)
			return
		}
		details, err := h.StoreService.GetProductDetails(product.ProductID)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		if product.ProductID != details.ProductID {
			response.JSON(context, http.StatusBadRequest, nil, nil, "please enter a valid product id")
			return
		}
		response.JSON(context, http.StatusCreated, details, nil, "products fetched successfully")
	}
}
