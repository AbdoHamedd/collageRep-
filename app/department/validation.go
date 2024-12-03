package department

import (
	"basicCrudoperations/exchanges"
	"basicCrudoperations/models"
	"basicCrudoperations/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func createDepartmentValidation(c *gin.Context) (error, exchanges.CreateDepartmentRequest) {
	var req exchanges.CreateDepartmentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, exchanges.CreateDepartmentRequest{}
	}
	err = validateNameAndDescription(req.Name, req.Description)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, exchanges.CreateDepartmentRequest{}
	}
	if !ValidateDepartmentNameIsUnique(req.Name) {
		err = errors.New("department should be uniqe")
		response.BadRequest(c, err.Error())
		return err, exchanges.CreateDepartmentRequest{}
	}
	return nil, req
}

func updateDepartmentValidation(c *gin.Context) (error, models.Department, exchanges.UpdateDepartmentRequest) {
	var request exchanges.UpdateDepartmentRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, models.Department{}, exchanges.UpdateDepartmentRequest{}
	}
	check, department := validateDepartmentIsFound(request.ID)
	{
		if !check {
			err = errors.New("department is not found")
			response.BadRequest(c, err.Error())
			return err, models.Department{}, exchanges.UpdateDepartmentRequest{}
		}
	}
	if !validateDepartmentNameIsUniqeForUpdate(request.Name, request.ID) {
		err = errors.New("department should be uniqe for update")
		response.BadRequest(c, err.Error())
		return err, models.Department{}, exchanges.UpdateDepartmentRequest{}
	}
	return nil, department, request

}

func validateNameAndDescription(Name string, Description string) error {
	var err error
	if len(Name) < 2 || len(Name) > 10 {
		err = errors.New(Name + " name must be between 2 and 10")
		return err
	}
	if len(Description) < 5 || len(Description) > 50 {
		err = errors.New(Description + " description must be between 5 and 50")
		return err
	}
	return nil
}

func deleteDepartmentValidation(c *gin.Context) (error, uint) {
	var req exchanges.DeleteDepartmentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, req.ID
	}
	check, _ := validateDepartmentIsFound(req.ID)
	if !check {
		err = errors.New("department is found")
		response.BadRequest(c, err.Error())
		return err, req.ID
	}
	return nil, req.ID
}

func getDepartmentByIdValidation(c *gin.Context) (error, models.Department) {
	var req exchanges.GetDepartmentByIdRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, models.Department{}
	}
	check, department := validateDepartmentIsFound(req.ID)
	if !check {
		err = errors.New("department is not found")

		response.BadRequest(c, err.Error())
		return err, models.Department{}
	}
	return nil, department
}

//func getAllDepartmentsValidation(c *gin.Context) error  {
//
//	err := errors.New()
//}
