package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// [POST] create
func CreateProduct(c *gin.Context) {
	c.String(http.StatusOK, "CreateProduct")
}

// [GET] all
func GetProducts(c *gin.Context) {
	c.String(http.StatusOK, "GetProducts")
}

// [GET] one
func GetProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	c.String(http.StatusOK, "GetProduct "+id)
}

// [PUT] update
func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	c.String(http.StatusOK, "UpdateProduct "+id)
}

// [DELETE] delete
func DeleteProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	c.String(http.StatusOK, "DeleteProduct "+id)
}
