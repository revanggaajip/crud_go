package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/revanggaajip/crud_go/models"
	"github.com/revanggaajip/crud_go/repository"
)

func CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindBodyWithJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := repository.CreateProduct(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
