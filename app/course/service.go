package course

import (
	"basicCrudoperations/exchanges/courseExchabge"
	"basicCrudoperations/models"
)

func CreateCourseService(request courseExchabge.CreateCourseRequest) models.Course {
	course := setCourse(request)
	createCourseModel(&course)
	return course
}
func UpdateCourseService(request courseExchabge.UpdateCourseRequest, courseBeforeUpdate models.Course) models.Course {
	courseUpdated := setCourseForUpdate(request, courseBeforeUpdate)
	UpdateCourseModel(&courseUpdated)
	return courseUpdated
}

func DeleteCourseService(id uint) {
	DeleteCourseModel(id)
}

func ShowAllCoursesServices(limit, offset int, order string) []models.Course {
	courses := ShowAllCourses(limit, offset, order)
	return courses
}
