package department

import (
	"basicCrudoperations/DataBase"
	"basicCrudoperations/models"
)

func createDepartmentModel(department *models.Department) {
	DataBase.DB.Preload("Courses").Create(&department)
}

func updateDepartmentModel(department *models.Department) {
	DataBase.DB.Preload("Courses").Save(&department)
}

func ValidateDepartmentNameIsUnique(name string) bool {
	var count int64
	DataBase.DB.Model(&models.Department{}).Where("name = ?", name).Count(&count)
	return count == 0
}

func validateDepartmentNameIsUniqeForUpdate(name string, id uint) bool {
	var count int64
	DataBase.DB.Model(&models.Department{}).Where("name = ?", name).Where("id != ?", id).Count(&count)
	return count == 0
}

func ValidateDepartmentIsFound(id uint) (bool, models.Department) {
	department := models.Department{}
	DataBase.DB.Preload("Courses").Where("id = ?", id).Find(&department)
	return department.ID != 0, department
}

func deleteDepartmentModel(id uint) {

	DataBase.DB.Model(&models.Department{}).Where("id = ?", id).Delete(&models.Department{})
}

func ShowAllDepartments(limit, offset int, order string) []models.Department {
	departments := []models.Department{}
	DataBase.DB.Preload("Courses").Limit(limit).Offset(offset).Order(order).Find(&departments)
	return departments
}

// new but not work correctky
func ValidateDepartmentIsUsed(id uint) bool {
	var count int64
	DataBase.DB.Model(&models.Course{}).Where("department_ID = ?", id).Count(&count)
	return count != 0
}
