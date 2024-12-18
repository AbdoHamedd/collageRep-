package department

import (
	"basicCrudoperations/exchanges/courseExchabge/CommonPagination"
	"basicCrudoperations/response"
	"basicCrudoperations/tranform"
	"github.com/gin-gonic/gin"
)

func createDepartment(c *gin.Context) {
	err, request := createDepartmentValidation(c)
	if err != nil {
		return
	}
	department := CreateDepartmentService(request)
	res := tranform.DepartmentTransFormReturn(department)
	response.Created(c, res)
}

func updateDepartment(c *gin.Context) {
	err, department, request := updateDepartmentValidation(c)
	if err != nil {
		return
	}
	department = updateDepartmentServices(request, department)
	res := tranform.DepartmentTransFormReturn(department)
	response.Ok(c, res)

}

func deleteDepartment(c *gin.Context) {
	err, id := deleteDepartmentValidation(c)
	if err != nil {
		return
	}
	deleteDepartmentService(id)
	response.Deleted(c)

}

func getDepartmentById(c *gin.Context) {
	err, department := getDepartmentByIdValidation(c)
	if err != nil {
		return
	}
	res := tranform.DepartmentTransFormReturn(department)
	response.Ok(c, res)
}

func getAllDepartments(c *gin.Context) {
	err, req := CommonPagination.ShawAllPaginationWithOrder(c)
	if err != nil {
		return
	}
	departments := showAll(req.Limit, req.Offset, req.Order)
	res := tranform.DepartmentsTransformFormReturn(departments)
	response.Ok(c, res)
}
