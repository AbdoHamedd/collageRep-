package models

import (
	"basicCrudoperations/DataBase"
)

func Migrate() {
	DataBase.DB.AutoMigrate(&Course{}) // as refernce
	DataBase.DB.AutoMigrate(&Department{})
}
