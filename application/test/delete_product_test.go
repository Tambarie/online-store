package test

import (
	"bytes"
	"encoding/json"
	"github.com/Tambarie/online-store/application/handler"
	"github.com/Tambarie/online-store/application/server"
	"github.com/Tambarie/online-store/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDeleteProduct(t *testing.T) {
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

	marhall, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	mockProductService.EXPECT().DeleteProduct(product.ProductID, gomock.Any()).Return(nil)

	t.Run("test for create product", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/api/v1/delete-product", bytes.NewBuffer(marhall))
		if err != nil {
			log.Fatalf("an error occured:%v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)

		if response.Code != http.StatusOK {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}
		var responseBodyTwo = `"product delete successfully"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}
	})
}
