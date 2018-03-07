package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/mihailo-misic/company-resource-api/controllers"
	"github.com/mihailo-misic/company-resource-api/database"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		// Users
		users := v1.Group("/users")
		{
			users.POST("/", c.CreateUser)
			users.GET("/", c.GetUsers)
			users.GET("/:id", c.GetUser)
			users.PUT("/:id", c.UpdateUser)
			users.DELETE("/:id", c.DeleteUser)
		}
		// Products
		products := v1.Group("/products")
		{
			products.POST("/", c.CreateProduct)
			products.GET("/", c.GetProducts)
			products.GET("/:id", c.GetProduct)
			products.PUT("/:id", c.UpdateProduct)
			products.DELETE("/:id", c.DeleteProduct)
		}
		// Auth
	}

	return r
}

func main() {
	// Setup database
	db := database.Init()
	defer db.Close()

	// Setup routing
	r := setupRouter()

	// Run the api on: localhost:8080
	r.Run(":8080")
}
