package product

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	productService := CreateProductService()
	ctrl := CreateProductController(productService)
	productGroup := router.Group("/products")
	{
		productGroup.GET("/:id", ctrl.FindProductById)
	}
}
