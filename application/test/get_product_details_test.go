package test

import (
	"bytes"
	"encoding/json"
	"github.com/Tambarie/online-store/application/handler"
	"github.com/Tambarie/online-store/application/server"
	domain "github.com/Tambarie/online-store/domain/store"
	"github.com/Tambarie/online-store/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetProductDetails(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := gomock.NewController(t)
	mockProductService := service.NewMockStoreService(controller)
	router := gin.Default()
	h := &handler.Handler{StoreService: mockProductService}
	server.DefineRouter(router, h)

	product := struct {
		ProductID string `json:"product_id"`
	}{
		ProductID: "1",
	}

	productDetails := &domain.Products{
		Model:       domain.Model{},
		ProductID:   "1",
		Name:        "mac-book",
		Description: "it's nice",
		Price:       120000.00,
		Quantity:    5,
		Status:      domain.INSTOCK,
	}

	marhall, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	mockProductService.EXPECT().GetProductDetails(product.ProductID).Return(productDetails, nil)

	t.Run("test for create product", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/get-product-details", bytes.NewBuffer(marhall))
		if err != nil {
			log.Fatalf("an error occured:%v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)

		log.Println(response.Body)

		if response.Code != http.StatusCreated {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}
		var responseBodyTwo = `"product_id":"1","name":"mac-book"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}
	})
}
