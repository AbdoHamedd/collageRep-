package department

import "github.com/gin-gonic/gin"

func DepartmentRoute(router *gin.Engine) {
	r := router.Group("/department")
	r.POST("/create", createDepartment)
	r.PUT("/update", updateDepartment)
	r.DELETE("/delete", deleteDepartment)
	r.GET("/getById", getDepartmentById)
	r.GET("/getAll", getAllDepartments)

}
