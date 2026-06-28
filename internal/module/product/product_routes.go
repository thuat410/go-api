package product

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(router *gin.RouterGroup, dbPool *pgxpool.Pool) {
	productService := CreateProductService(dbPool)
	ctrl := CreateProductController(productService)
	productGroup := router.Group("/products")
	{
		productGroup.GET("/:id", ctrl.FindProductById)
	}
}
