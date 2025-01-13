package student_course

import (
	"basicCrudoperations/app/course"
	"basicCrudoperations/app/user"
	"basicCrudoperations/exchanges/studenr_course"
	"basicCrudoperations/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func StudentJoinCourseValidate(c *gin.Context) bool {
	var req studenr_course.StudentJoinCourseRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return false
	}

	check, _ := course.ValidateCourseIsFound(req.CourseID)
	if !check {
		err = errors.New("Course  is not Found")

		response.BadRequest(c, err.Error())
		return false
	}
	check, _ = user.UserIsFound(req.UserID)
	if !check {
		err = errors.New("user is not Found")

		response.BadRequest(c, err.Error())
		return false
	}
	check = userIsStudent(req.UserID)
	if !check {
		err = errors.New("user is not Student")
		response.BadRequest(c, err.Error())
		return false
	}
	check = joinValidate(req)
	if !check {
		err = errors.New("you are already Register")
		response.BadRequest(c, err.Error())
		return false
	}
	return true
}

func StudendDrobCourseValidate(c *gin.Context) bool {
	var req studenr_course.StudentJoinCourseRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return false
	}
	check := joinValidate(req)
	if check {
		err = errors.New("can't Found This Register")
		response.BadRequest(c, err.Error())
		return false
	}
	return true
}
