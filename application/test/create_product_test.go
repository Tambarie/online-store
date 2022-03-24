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

func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := gomock.NewController(t)
	mockProductService := service.NewMockStoreService(controller)
	router := gin.Default()
	h := &handler.Handler{StoreService: mockProductService}
	server.DefineRouter(router, h)

	product := &domain.Products{
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
	mockProductService.EXPECT().CreateProduct(gomock.Any()).Return(product, nil)

	t.Run("test for create product", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/api/v1/create-product", bytes.NewBuffer(marhall))
		if err != nil {
			log.Fatalf("an error occured:%v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)

		if response.Code != http.StatusCreated {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}
		var responseBodyTwo = `"message":"product created successfully"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}
	})
}
