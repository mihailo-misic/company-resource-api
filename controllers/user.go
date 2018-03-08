package controllers

import (
	"github.com/gin-gonic/gin"
	m "github.com/mihailo-misic/company-resource-api/models"
	"net/http"
	"github.com/mihailo-misic/company-resource-api/res"
)

// [POST] create
func CreateUser(c *gin.Context) {
	var user m.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, res.Err(res.Error{"Could not bind to JSON", "Error occured while binding the data to JSON", err}))
		return
	}

	db.Save(&user)

	if user.ID != 0 {
		c.JSON(http.StatusOK, res.Data(user))
		return
	}
	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Email taken", "The chosen email is already taken", nil}))
}

// [GET] all
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, res.Data(db.Find(&[]m.User{})))
}

// [GET] one
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	c.JSON(http.StatusOK, res.Data(db.First(&m.User{}, id)))
}

// [PUT] update
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user m.User
	var newUser m.ResUser

	// Find the user
	db.First(&user, id)
	// Parse the request data
	c.BindJSON(&newUser)
	// Update the user in the database
	db.Model(&user).Updates(newUser)

	c.JSON(http.StatusOK, res.Data(user))
}

// [DELETE] delete
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user m.User

	db.Delete(&user, id)

	if user.ID != 0 {
		c.JSON(http.StatusOK, res.Data(user))
	}

	c.JSON(http.StatusBadRequest, res.Err(res.Error{"Email taken", "The chosen email is already taken", nil}))
}
