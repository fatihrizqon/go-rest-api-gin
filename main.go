package main

import (
	"github.com/fatihrizqon/go-rest-api-gin/controllers"
	"github.com/fatihrizqon/go-rest-api-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/api/products", controllers.GetProducts)
	r.POST("/api/products", controllers.CreateProduct)
	r.GET("/api/products/:id", controllers.GetProductById)
	r.PUT("/api/products/:id", controllers.UpdateProductById)
	r.DELETE("/api/products/:id", controllers.DeleteProductById)

	r.Run()
}
