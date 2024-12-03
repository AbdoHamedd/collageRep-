package department

import (
	"basicCrudoperations/exchanges"
	"basicCrudoperations/models"
)

func SetDepartment(request exchanges.CreateDepartmentRequest) models.Department {
	return models.Department{
		Name:        request.Name,
		Description: request.Description,
	}
}

func setUpdatedDepartment(req exchanges.UpdateDepartmentRequest, department models.Department) models.Department {
	department.Name = req.Name
	department.Description = req.Description
	return department
}
