package server

import (
	"fmt"
	"github.com/Tambarie/online-store/application/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func initializeRouter() *gin.Engine {
	router := gin.Default()
	if os.Getenv("GIN_MODE") == "testing" {
		return router
	}

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	// setup cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	return router
}

func DefineRouter(router *gin.Engine, handler *handler.Handler) {
	apiRouter := router.Group("/api/v1")
	apiRouter.POST("/create-product", handler.CreateProduct())
	apiRouter.GET("/get-product-details", handler.GetProductDetails())
	apiRouter.PUT("/update-product/:product_id", handler.UpdateProducts())
	apiRouter.DELETE("/delete-product", handler.DeleteProduct())
	apiRouter.GET("/list-all-products", handler.ListAllProducts())
	apiRouter.GET("/filter-products", handler.FilterProducts())
	apiRouter.POST("/add-products-basket", handler.AddProductToBasket())
	apiRouter.GET("/get-summary", handler.GetSummaryInBasket())
}
