package department

import (
	"basicCrudoperations/exchanges/deptexchange"
	"basicCrudoperations/models"
)

func SetDepartment(request deptexchange.CreateDepartmentRequest) models.Department {
	return models.Department{
		Name:        request.Name,
		Description: request.Description,
	}
}

func setUpdatedDepartment(req deptexchange.UpdateDepartmentRequest, department models.Department) models.Department {
	department.Name = req.Name
	department.Description = req.Description
	return department
}
