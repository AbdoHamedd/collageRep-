package tranform

import (
	"basicCrudoperations/exchanges/deptexchange"
	"basicCrudoperations/exchanges/shared"
	"basicCrudoperations/models"
)

func DepartmentTransform(department models.Department) deptexchange.DepartmentResponse {
	return deptexchange.DepartmentResponse{
		ID:          department.ID,
		CreatedAt:   department.CreatedAt,
		UpdatedAt:   department.UpdatedAt,
		Name:        department.Name,
		Description: department.Description,
	}
}

// new
func DepartmentTransFormReturn(department models.Department) shared.DepartmentResponseForDept {
	return shared.DepartmentResponseForDept{
		ID:          department.ID,
		CreatedAt:   department.CreatedAt,
		UpdatedAt:   department.UpdatedAt,
		Name:        department.Name,
		Description: department.Description,
		Courses:     CoursesTransformForReturnDepartment(department.Courses),
	}
}

func DepartmentsTransformFormReturn(departments []models.Department) (response []shared.DepartmentResponseForDept) {
	for _, department := range departments {
		response = append(response, DepartmentTransFormReturn(department))
	}
	return response
}
