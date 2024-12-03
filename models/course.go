package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	TotalHours   int        `json:"totalHours"`
	Department   Department `json:"department"`
	DepartmentID uint64     `json:"departmentID" validate:"required"`
}
