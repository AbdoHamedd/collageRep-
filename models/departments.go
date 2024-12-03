package models

import (
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name        string   `json:"name" validate:"required" gorm:"type:varchar(50);unique"`
	Description string   `json:"description" validate:"required"  gorm:"type:varchar(500)"`
	courses     []Course `json:"course"`
}
