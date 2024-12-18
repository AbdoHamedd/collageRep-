package course

import (
	"basicCrudoperations/DataBase"
	"basicCrudoperations/models"
)

func createCourseModel(course *models.Course) {
	DataBase.DB.Preload("Department").Create(&course)
}

func UpdateCourseModel(course *models.Course) {
	DataBase.DB.Preload("Department").Save(&course)
}

func validateNameIsUnique(name string) bool {
	var count int64
	DataBase.DB.Model(&models.Course{}).Where("name = ?", name).Count(&count)
	return count == 0
}

func ValidateCourseIsFound(id uint) (bool, models.Course) {
	course := models.Course{}
	DataBase.DB.Preload("Department").Where("id = ?", id).Find(&course)
	return course.ID != 0, course
}

func validateNameIsUniqueForUpdate(name string, id uint) bool {
	var count int64
	DataBase.DB.Model(&models.Course{}).Where("name = ?", name).Where("id != ?", id).Count(&count)
	return count == 0
}

func DeleteCourseModel(id uint) {
	DataBase.DB.Where("id =?", id).Delete(&models.Course{})
}

func ShowAllCourses(limit, offset int, order string) []models.Course {
	courses := []models.Course{}
	DataBase.DB.Preload("Department").Limit(limit).Offset(offset).Order(order).Find(&courses)
	return courses
}

func ValidateDepartmentIsFound(id uint) (bool, models.Department) {
	department := models.Department{}
	DataBase.DB.Model(&models.Department{}).Where("id = ?", id).Find(&department)
	return department.ID != 0, department
}
