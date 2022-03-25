package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	domain "github.com/Tambarie/online-store/domain/store"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (h *Handler) CreateProduct() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Logging the context
		helpers.LogRequest(context)

		product := &domain.Products{}

		product.ProductID = uuid.New().String()
		product.Status = domain.INSTOCK

		// Binding the json
		if err := helpers.Decode(context, &product); err != nil {
			log.Fatalf("Error %v", err)
			return
		}

		// Creates a product
		productDB, err := h.StoreService.CreateProduct(product)
		if err != nil {
			context.AbortWithStatusJSON(500, err)
			return
		}
		// JSON Response to the client
		response.JSON(context, http.StatusCreated, productDB.ProductID, nil, "product created successfully")
	}
}
