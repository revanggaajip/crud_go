package repository

import (
	"log"

	"github.com/revanggaajip/crud_go/database"
	"github.com/revanggaajip/crud_go/models"
)

func CreateProduct(product *models.Product) error {
	if err := database.DB.Create(product).Error; err != nil {
		log.Println("Error creating product:", err)
		return err
	}
	return nil
}
