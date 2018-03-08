package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/mihailo-misic/company-resource-api/controllers"
	"github.com/mihailo-misic/company-resource-api/database"
	"os"
	"io"
	"log"
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
	var f *os.File
	// Setup logging
	f, _ = os.OpenFile("crm.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	// Setup database
	db := database.Init()
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	// Setup routing
	r := setupRouter()

	// Run the api on: "localhost:8080"
	log.Fatal(r.Run(":8080"))
}
