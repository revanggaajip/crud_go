package main

import (
	"github.com/gin-gonic/gin"
	"github.com/revanggaajip/crud_go/database"
	"github.com/revanggaajip/crud_go/routes"
)

func main() {

	// Inisialisasi database
	database.InitDatabase()

	// Menjalankan migrate
	database.InitMigrate()

	// Inisialisasi router
	router := gin.Default()
	// Mengimpor route dari package routes
	routes.RegisterRoutes(router)

	router.Run(":5000")
}
