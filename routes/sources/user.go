package sources

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User Route",
			})
		})

		userGroup.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User ID: " + id,
			})
		})
	}
}
