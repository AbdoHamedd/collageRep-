package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"varchar(50)"`
	Email    string `json:"email" validate:"required" gorm:"varchar(50)"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" gorm:"type:varchar(10);default:'Student'"`
	Token    string `json:"token"`
	IsAusth  bool
}
