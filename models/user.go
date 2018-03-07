package models

import "github.com/jinzhu/gorm"

type (
	// User describes a User type
	User struct {
		gorm.Model
		Email     string `json:"email" gorm:"unique_index;not null;type:varchar(100);" binding:"required"`
		Password  string `json:"password" binding:"required"`
		FirstName string `json:"firstName" gorm:"column:first_name" binding:"required"`
		LastName  string `json:"lastName" gorm:"column:last_name" binding:"required"`
		Image     string `json:"image"`
	}
	// ResUser represents a formatted User
	ResUser struct {
		ID        uint   `json:"id"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Image     string `json:"image"`
	}
)
