package models

import "gorm.io/gorm"

type StusentCourse struct {
	gorm.Model
	UserId   uint   `json:"user_id" validate:"requires"`
	CourseId uint   `json:"course_id" validate:"required"`
	User     User   `json:"user"`
	Course   Course `json:"course"`
}
