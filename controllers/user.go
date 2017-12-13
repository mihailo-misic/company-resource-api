package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/mihailo-misic/company-resource-api/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int64} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	id := models.AddUser(user)
	u.Data["json"] = map[string]int64{"id": id}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (u *UserController) Get() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.Data["json"] = err.Error()
	}

	if id != 0 {
		user, err := models.GetUser(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}

	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	id		path 	int64     	true		"The id you want to update"
// @Param	body	body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (u *UserController) Put() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.Data["json"] = err.Error()
	}

	if id != 0 {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(id, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	id		path 	int64	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (u *UserController) Delete() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.Data["json"] = err.Error()
	}

	models.DeleteUser(id)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	email		query 	string	true		"The email for login"
// @Param	password	query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	if models.Login(user.Email, user.Password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user does not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [post]
func (u *UserController) Logout() {
	if models.Logout() {
		u.Data["json"] = "logout success"
	} else {
		u.Data["json"] = "logout success"
	}
	u.ServeJSON()
}
