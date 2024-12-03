package course

import (
	"basicCrudoperations/exchanges"
	"basicCrudoperations/models"
	"basicCrudoperations/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func CreateCourseValidation(c *gin.Context) (error, exchanges.CreateCourseRequest) {
	var request exchanges.CreateCourseRequest
	err := c.ShouldBindJSON(&request) // validate request is json ,  set request
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, exchanges.CreateCourseRequest{}
	}
	err = ValidateCourseNameDescriptionTotalHours(request.Name, request.Description, request.TotalHours)
	if err != nil {
		response.BadRequest(c, err.Error())

		return err, exchanges.CreateCourseRequest{}
	}
	if !validateNameIsUnique(request.Name) {
		err = errors.New("Name should be unique")
		response.BadRequest(c, err.Error())
		return err, exchanges.CreateCourseRequest{}
	}
	return nil, request
}

func updateCourseValidation(c *gin.Context) (error, exchanges.UpdateCourseRequest, models.Course) {
	var request exchanges.UpdateCourseRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, exchanges.UpdateCourseRequest{}, models.Course{}
	}
	check, course := ValdateCourseIsFound(request.ID)
	if !check {
		err = errors.New("Course ID is not Found")
		response.BadRequest(c, err.Error())
		return err, exchanges.UpdateCourseRequest{}, models.Course{}
	}
	if !validateNameIsUniqueForUpdate(request.Name, request.ID) {
		err = errors.New("Name should be unique")
		response.BadRequest(c, err.Error())
		return err, exchanges.UpdateCourseRequest{}, models.Course{}
	}
	err = ValidateCourseNameDescriptionTotalHours(request.Name, request.Description, request.TotalHours)
	if err != nil {
		response.BadRequest(c, err.Error())

		return err, exchanges.UpdateCourseRequest{}, models.Course{}
	}
	return nil, request, course

}

func ValidateCourseNameDescriptionTotalHours(name string, description string, totalHours int) error {
	var err error
	if len(name) < 2 || len(name) > 50 {
		err = errors.New("Name must be at between 2 and 50 characters")
		return err
	}
	if len(description) < 2 || len(description) > 500 {
		err = errors.New("Description must be at between 2 and 500 characters")
		return err
	}
	if totalHours != 0 && totalHours != 2 && totalHours != 3 {
		err = errors.New("Total hours must be 0 or 2 or 3")
		return err
	}
	return nil
}

func deleteCourseValidations(c *gin.Context) (error, uint) {
	var request exchanges.DeleteCourseRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.BadRequest(c, err.Error())
		return err, request.ID
	}
	// Check if the course exists
	check, _ := ValdateCourseIsFound(request.ID)
	{
		if !check {
			err = errors.New("Course  is not Found")

			response.BadRequest(c, err.Error())
			return err, request.ID
		}
	}
	return nil, request.ID
}

func getCourseByIdValidation(c *gin.Context) (error, models.Course) {
	var req exchanges.GetCourseByIdRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return nil, models.Course{}
	}
	check, Course := ValdateCourseIsFound(req.ID)

	if !check {
		err = errors.New("Course  is not Found")
		response.BadRequest(c, err.Error())
		return err, models.Course{}
	}

	return nil, Course
}

//func getAllCoursesValidation(c *gin.Context) error {
//	var req exchanges.GetCourseByIdRequest
//	err := c.ShouldBindJSON(&req)
//	if err != nil {
//		response.BadRequest(c, err.Error())
//		return err
//	}
//	return nil
//}
