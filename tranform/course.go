package tranform

import (
	"basicCrudoperations/exchanges/courseExchabge"
	"basicCrudoperations/exchanges/shared"
	"basicCrudoperations/models"
	"time"
)

func CourseTransform(course models.Course) shared.CourseResponse {
	return shared.CourseResponse{
		ID:          course.ID,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   time.Time{},
		Name:        course.Name,
		Description: course.Description,
		TotalHours:  course.TotalHours,
		Department:  DepartmentTransform(course.Department),
	}
}

func CoursesTransform(courses []models.Course) (response []shared.CourseResponse) {
	for _, course := range courses {
		response = append(response, CourseTransform(course))
	}
	return
}

//-----------------New---------------------------------------

func CourseTransformForReturnDepartment(course models.Course) courseExchabge.CourseResoForDept {
	return courseExchabge.CourseResoForDept{
		ID:          course.ID,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   time.Time{},
		Name:        course.Name,
		Description: course.Description,
		TotalHours:  course.TotalHours,
	}
}

func CoursesTransformForReturnDepartment(courses []models.Course) (response []courseExchabge.CourseResoForDept) {
	for _, course := range courses {
		response = append(response, CourseTransformForReturnDepartment(course))
	}
	return
}
