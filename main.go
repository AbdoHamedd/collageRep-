package main

import (
	"basicCrudoperations/DataBase"
	"basicCrudoperations/models"
	"basicCrudoperations/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // get default engin to use to make routes

	DataBase.ConnectToDataBase() // connection with logger

	models.Migrate() // create tables (models) // make once but should be run if add new model or update any model

	routes.Routes(r) // COLLECT ALL curd(restful apis) from all modules

	r.Run(":5050") // run engin with port 5050

}
