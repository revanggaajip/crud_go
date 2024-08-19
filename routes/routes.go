package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/revanggaajip/crud_go/routes/sources"
)

func RegisterRoutes(router *gin.Engine) {
	sources.UserRoutes(router)
	sources.ProductRoutes(router)
}
