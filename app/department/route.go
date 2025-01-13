package department

import (
	"basicCrudoperations/auth"
	"github.com/gin-gonic/gin"
)

func DepartmentRoute(router *gin.Engine) {
	r := router.Group("/department").Use(auth.Auth).Use(auth.AuthAdmin)
	r.POST("/create", createDepartment)
	r.PUT("/update", updateDepartment)
	r.DELETE("/delete", deleteDepartment)
	r.GET("/getById", getDepartmentById)
	r.GET("/getAll", getAllDepartments)

}
