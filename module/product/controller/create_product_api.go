package productcontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	productdomain "myapp/module/product/domain"
	"net/http"
)

func (api APIController) CreateProductAPI(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var productData productdomain.ProductCreateDTO

		if err := c.Bind(&productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := api.createUseCase.CreateProduct(c.Request.Context(), &productData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": productData.Id})
	}
}
