package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int `orm:"pk"`
	Email     string
	Password  string
	FirstName string
	LastName  string
	Image     string `orm:"null"`
	UpdatedAt string
	CreatedAt string
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(usr User) (id int) {
	// TODO
	return usr.Id
}

func GetUser(id string) (usr *User, err error) {
	// TODO
	return nil, errors.New("user not exists")
}

func GetAllUsers() map[string]*User {
	// TODO
	return make(map[string]*User)
}

func UpdateUser(id string, usr *User) (updUsr *User, err error) {
	// TODO
	return nil, errors.New("user Not Exist")
}

func Login(username, password string) (suc bool) {
	// TODO
	return false
}

func DeleteUser(id string) (err error) {
	// TODO
	return nil
}
