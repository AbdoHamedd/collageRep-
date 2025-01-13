package user

import (
	"basicCrudoperations/DataBase"
	"basicCrudoperations/models"
)

func createModelUser(user *models.User) {
	DataBase.DB.Create(&user)
}

func emailIsExists(email string) (bool, models.User) {
	user := models.User{}
	DataBase.DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return false, models.User{}
	}
	return true, user
}
func updateUser(user *models.User) {
	DataBase.DB.Save(&user)
}
func UserIsFound(id uint) (bool, models.User) {
	user := GetUserById(id)
	return user.ID != 0, user
}

func GetUserById(id uint) models.User {
	user := models.User{}
	DataBase.DB.Where("id = ?", id).First(&user)
	return user
}
