package department

import (
	"basicCrudoperations/exchanges/deptexchange"
	"basicCrudoperations/models"
)

func CreateDepartmentService(req deptexchange.CreateDepartmentRequest) models.Department {
	department := SetDepartment(req)
	createDepartmentModel(&department)
	return department
}

func updateDepartmentServices(req deptexchange.UpdateDepartmentRequest, departmet models.Department) models.Department {
	departmet = setUpdatedDepartment(req, departmet)
	updateDepartmentModel(&departmet)
	return departmet
}
func deleteDepartmentService(id uint) {
	deleteDepartmentModel(id)
}

func showAll(limit, Offset int, order string) []models.Department {
	departments := ShowAllDepartments(limit, Offset, order)
	return departments
}
