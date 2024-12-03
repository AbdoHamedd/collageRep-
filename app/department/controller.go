package department

import (
	"basicCrudoperations/response"
	"github.com/gin-gonic/gin"
)

func createDepartment(c *gin.Context) {
	err, request := createDepartmentValidation(c)
	if err != nil {
		return
	}
	department := CreateDepartmentService(request)
	response.Created(c, department)
}

func updateDepartment(c *gin.Context) {
	err, department, request := updateDepartmentValidation(c)
	if err != nil {
		return
	}
	updateDepartmentServices(request, department)
	response.Ok(c, department)

}
