package controllers

import (
	"github.com/gin-gonic/gin"
	m "github.com/mihailo-misic/company-resource-api/models"
	"net/http"
)

// [POST] create
func CreateUser(c *gin.Context) {
	var user m.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db.Save(&user)

	if user.ID != 0 {
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already taken"})
}

// [GET] all
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, db.Find(&m.User{}))
}

// [GET] one
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, db.First(&m.User{}, id))
}

// [PUT] update
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user m.User

	db.First(&user, id)

	//TODO continue

	c.JSON(http.StatusOK, "UpdateUser "+id)
}

// [DELETE] delete
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, "DeleteUser "+id)
}
