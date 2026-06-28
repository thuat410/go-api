package product

import (
	"go-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *ProductService
}

func CreateProductController(s *ProductService) *ProductController {
	return &ProductController{productService: s}
}

func (ctrl *ProductController) FindProductById(c *gin.Context) {
	productID := c.Param("id")
	productData, err := ctrl.productService.FindProductById(productID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	utils.OKResponse(c, productData)
}
