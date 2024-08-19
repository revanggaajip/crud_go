package database

import (
	"fmt"
	"log"

	"github.com/revanggaajip/crud_go/models"
)

func InitMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database migrated")

}
