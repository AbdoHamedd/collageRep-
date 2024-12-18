package course

import (
	"basicCrudoperations/exchanges/courseExchabge"
	"basicCrudoperations/models"
)

func setCourse(request courseExchabge.CreateCourseRequest) models.Course {
	return models.Course{
		Name:         request.Name,
		Description:  request.Description,
		TotalHours:   request.TotalHours,
		DepartmentID: request.DepartmentID,
	}
}

func setCourseForUpdate(request courseExchabge.UpdateCourseRequest, course models.Course) models.Course {
	course.Name = request.Name
	course.Description = request.Description
	course.TotalHours = request.TotalHours
	course.DepartmentID = request.DepartmentID
	return course
}
