package course

import (
	"basicCrudoperations/exchanges"
	"basicCrudoperations/models"
)

func setCourse(request exchanges.CreateCourseRequest) models.Course {
	return models.Course{
		Name:        request.Name,
		Description: request.Description,
		TotalHours:  request.TotalHours,
	}
}

func setCourseForUpdate(request exchanges.UpdateCourseRequest, course models.Course) models.Course {
	course.Name = request.Name
	course.Description = request.Description
	course.TotalHours = request.TotalHours
	return course
}
