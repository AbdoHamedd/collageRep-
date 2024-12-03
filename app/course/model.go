package course

import (
	"basicCrudoperations/DataBase"
	"basicCrudoperations/models"
)

func createCourseModel(course *models.Course) {
	DataBase.DB.Create(&course)
}

func UpdateCourseModel(course *models.Course) {
	DataBase.DB.Save(&course)
}

func validateNameIsUnique(name string) bool {
	var count int64
	DataBase.DB.Model(&models.Course{}).Where("name = ?", name).Count(&count)
	return count == 0
}

func ValdateCourseIsFound(id uint) (bool, models.Course) {
	course := models.Course{}
	DataBase.DB.Model(&models.Course{}).Where("id = ?", id).Find(&course)
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

func ShowAllCourses() []models.Course {
	courses := []models.Course{}
	DataBase.DB.Find(&courses)
	return courses
}
