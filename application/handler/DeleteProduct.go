package handler

import (
	"github.com/Tambarie/online-store/application/helpers"
	domain "github.com/Tambarie/online-store/domain/store"
	"github.com/Tambarie/online-store/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) DeleteProduct() gin.HandlerFunc {
	return func(context *gin.Context) {
		product := struct {
			ProductID string `json:"product_id"`
		}{}

		production := &domain.Products{}

		if err := helpers.Decode(context, &product); err != nil {
			log.Fatalf("Error  cannot bind %v", err)
			return
		}

		err := h.StoreService.DeleteProduct(product.ProductID, production)
		if err != nil {
			return
		}

		response.JSON(context, http.StatusOK, nil, nil, "product delete successfully")

	}
}
