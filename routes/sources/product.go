package sources

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revanggaajip/crud_go/controllers"
)

func ProductRoutes(router *gin.Engine) {
	productGroup := router.Group("/product")
	{
		productGroup.POST("/", controllers.CreateProduct)
		productGroup.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "product route",
			})
		})

		productGroup.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Product ID" + id,
			})
		})
	}
}
