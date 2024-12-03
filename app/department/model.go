package department

import (
	"basicCrudoperations/DataBase"
	"basicCrudoperations/models"
)

func createDepartmentModel(department *models.Department) {
	DataBase.DB.Create(&department)
}

func updateDepartmentModel(department *models.Department) {
	DataBase.DB.Save(&department)
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
func validateDepartmentIsFound(id uint) (bool, models.Department) {
	department := models.Department{}
	DataBase.DB.Model(&models.Department{}).Where("id = ?", id).Find(&department)
	return department.ID != 0, department
}
