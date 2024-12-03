package department

import "github.com/gin-gonic/gin"

func DepartmentRoute(router *gin.Engine) {
	r := router.Group("/department")
	r.POST("/createDepartment", createDepartment)
	r.PUT("/updateDepartment", updateDepartment)
	r.DELETE("/deleteDepartment", deleteDepartment)
	r.GET(("/getDepartmentById"), getDepartmentById)

}
