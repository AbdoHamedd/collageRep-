package department

import (
	"basicCrudoperations/exchanges"
	"basicCrudoperations/models"
)

func CreateDepartmentService(req exchanges.CreateDepartmentRequest) models.Department {
	department := SetDepartment(req)
	createDepartmentModel(&department)
	return department
}

func updateDepartmentServices(req exchanges.UpdateDepartmentRequest, departmet models.Department) models.Department {
	updatedDepartment := setUpdatedDepartment(req, departmet)
	updateDepartmentModel(&updatedDepartment)
	return updatedDepartment
}
func deleteDepartmentService(id uint) {
	deleteDepartmentModel(id)
}
