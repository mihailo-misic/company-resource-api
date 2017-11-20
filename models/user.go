package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int64 `orm:"pk;auto"`
	Email     string
	Password  string
	FirstName string
	LastName  string
	Image     string    `orm:"null"`
	UpdatedAt time.Time `orm:"auto_now"`
	CreatedAt time.Time `orm:"auto_now_add"`
}

var o orm.Ormer

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(user User) (id int64) {
	o = orm.NewOrm()

	cryptPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(cryptPass[:])

	id, err := o.Insert(&user)
	if err != nil {
		fmt.Errorf("error occured creating a new user: %v", err)
	}

	return id
}

func GetUser(id int64) (user User, err error) {
	o = orm.NewOrm()
	user = User{Id: id}
	o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Printf("%+v\n", user)
		return user, nil
	}

	return User{}, errors.New("user not exists")
}

func GetAllUsers() (users []*User) {
	o = orm.NewOrm()
	num, err := o.QueryTable("user").All(&users)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	return users
}

func UpdateUser(id int64, user *User) (*User, error) {
	o = orm.NewOrm()
	oldUser := User{Id: id}
	err := o.Read(&oldUser)
	if err == nil {

		if user.Email != "" {
			oldUser.Email = user.Email
		}
		if user.Password != "" {
			cryptPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				return nil, errors.New("error changing the password")
			}
			oldUser.Password = string(cryptPass[:])
		}
		if user.FirstName != "" {
			oldUser.FirstName = user.FirstName
		}
		if user.LastName != "" {
			oldUser.LastName = user.LastName
		}
		if user.Image != "" {
			oldUser.Image = user.Image
		}

		if num, err := o.Update(&oldUser); err == nil {
			fmt.Println(num)

			return &oldUser, nil
		}
	}

	return nil, errors.New("user Not Exist")
}

func DeleteUser(id int64) (err error) {
	o = orm.NewOrm()
	num, err := o.Delete(&User{Id: id})
	if err != nil {
		return err
	}

	fmt.Println(num)
	return nil
}

func Login(email, password string) (suc bool) {
	o = orm.NewOrm()
	// TODO
	return false
}

func Logout() (suc bool) {
	// TODO
	return false
}
