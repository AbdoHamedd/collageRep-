package course

import (
	"basicCrudoperations/exchanges"
	"basicCrudoperations/models"
)

func CreateCourseService(request exchanges.CreateCourseRequest) models.Course {
	course := setCourse(request)
	createCourseModel(&course)
	return course
}
func UpdateCourseService(request exchanges.UpdateCourseRequest, courseBeforeUpdate models.Course) models.Course {
	courseUpdated := setCourseForUpdate(request, courseBeforeUpdate)
	UpdateCourseModel(&courseUpdated)
	return courseUpdated
}

func DeleteCourseService(id uint) {
	DeleteCourseModel(id)
}

func ShowAllCoursesServices() []models.Course {
	courses := ShowAllCourses()
	return courses
}
