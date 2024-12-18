package department

import (
	"basicCrudoperations/exchanges/deptexchange"
	"basicCrudoperations/models"
	"basicCrudoperations/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func createDepartmentValidation(c *gin.Context) (error, deptexchange.CreateDepartmentRequest) {
	var req deptexchange.CreateDepartmentRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, deptexchange.CreateDepartmentRequest{}
	}
	err = validateNameAndDescription(req.Name, req.Description)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, deptexchange.CreateDepartmentRequest{}
	}
	if !ValidateDepartmentNameIsUnique(req.Name) {
		err = errors.New("department should be uniqe")
		response.BadRequest(c, err.Error())
		return err, deptexchange.CreateDepartmentRequest{}
	}
	return nil, req
}

func updateDepartmentValidation(c *gin.Context) (error, models.Department, deptexchange.UpdateDepartmentRequest) {
	var request deptexchange.UpdateDepartmentRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, models.Department{}, deptexchange.UpdateDepartmentRequest{}
	}
	check, department := ValidateDepartmentIsFound(request.ID)

	if !check {
		err = errors.New("department is not found")
		response.BadRequest(c, err.Error())
		return err, models.Department{}, deptexchange.UpdateDepartmentRequest{}
	}

	if !validateDepartmentNameIsUniqeForUpdate(request.Name, request.ID) {
		err = errors.New("department should be uniqe for update")
		response.BadRequest(c, err.Error())
		return err, models.Department{}, deptexchange.UpdateDepartmentRequest{}
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

func getDepartmentByIdValidation(c *gin.Context) (error, models.Department) {
	var req deptexchange.GetByIdRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, models.Department{}
	}
	check, department := ValidateDepartmentIsFound(req.ID)
	if !check {
		err = errors.New("department is not found")

		response.BadRequest(c, err.Error())
		return err, models.Department{}
	}
	return nil, department
}

func deleteDepartmentValidation(c *gin.Context) (error, uint) {
	var req deptexchange.GetByIdRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, req.ID
	}
	check, _ := ValidateDepartmentIsFound(req.ID)
	if !check {
		err = errors.New("department is found")
		response.BadRequest(c, err.Error())
		return err, req.ID
	}
	//new
	check = ValidateDepartmentIsUsed(req.ID)
	if check {
		err = errors.New("department is already used")
		response.BadRequest(c, err.Error())
		return err, req.ID
	}
	return nil, req.ID
}

//func ValidateShawAll(c *gin.Context) (error, CommonPagination.Pagination) {
//	var req CommonPagination.Pagination
//	err := c.ShouldBindJSON(&req)
//	if err != nil {
//		response.BadRequest(c, err.Error())
//		return nil, CommonPagination.Pagination{}
//	}
//	CommonPagination.ApplyPaginationDefaults(&req)
//	return err, req
//}
