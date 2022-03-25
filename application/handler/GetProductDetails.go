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

		// Logging the context
		helpers.LogRequest(context)

		product := struct {
			ProductID string `json:"product_id"`
		}{}

		// Binding the json
		if err := helpers.Decode(context, &product); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		// Gets the product by it's ID from the database
		details, err := h.StoreService.GetProductDetails(product.ProductID)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		// Checks if the productID is valid
		if product.ProductID != details.ProductID {
			response.JSON(context, http.StatusBadRequest, nil, nil, "please enter a valid product id")
			return
		}
		// JSON Response to the client
		response.JSON(context, http.StatusCreated, details, nil, "products fetched successfully")
	}
}
