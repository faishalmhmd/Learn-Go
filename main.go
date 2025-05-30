package main

import (
	"rest-api-gin/controllers/productController"
	"rest-api-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectToDatabase()

	r.GET("/api/product", productController.Index)
	r.GET("/api/product/:id", productController.ShowById)
	r.POST("/api/product", productController.Create)
	r.PUT("/api/product", productController.Update)
	r.DELETE("/api/product", productController.Delete)

	r.Run()
}