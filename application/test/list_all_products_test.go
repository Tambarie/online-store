package test

import (
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

func TestListAllProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := gomock.NewController(t)
	mockProductService := service.NewMockStoreService(controller)
	router := gin.Default()
	h := &handler.Handler{StoreService: mockProductService}
	server.DefineRouter(router, h)

	product := []domain.Products{
		{
			Model:       domain.Model{},
			ProductID:   "1",
			Name:        "mac-book",
			Description: "it's nice",
			Price:       120000.00,
			Quantity:    5,
			Status:      domain.INSTOCK,
		},
		{
			Model:       domain.Model{},
			ProductID:   "2",
			Name:        "samsung tablet",
			Description: "it's nice",
			Price:       300000.00,
			Quantity:    2,
			Status:      domain.INSTOCK,
		},
	}
	//
	//marhall, err := json.Marshal(product)
	//if err != nil {
	//	log.Fatal(err)
	//}
	mockProductService.EXPECT().ListAllProducts().Return(product, nil)

	t.Run("test for create product", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/api/v1/list-all-products", nil)
		if err != nil {
			log.Fatalf("an error occured:%v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, req)

		if response.Code != http.StatusCreated {
			t.Errorf("expected %v but got %v", http.StatusCreated, response.Code)
		}
		var responseBodyTwo = `"description":"it's nice"`
		if !strings.Contains(response.Body.String(), responseBodyTwo) {
			t.Errorf("Expected body to contain %s", responseBodyTwo)
		}
	})
}
