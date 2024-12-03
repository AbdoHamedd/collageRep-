package course

import (
	"basicCrudoperations/response"
	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	err, request := CreateCourseValidation(c)
	if err != nil {
		return
	}
	course := CreateCourseService(request)
	response.Created(c, course)
}

func updateCourse(c *gin.Context) {
	err, request, course := updateCourseValidation(c)
	if err != nil {
		return
	}
	course = UpdateCourseService(request, course)
	response.Ok(c, course)
}

func deleteCourse(c *gin.Context) {
	err, ID := deleteCourseValidations(c)
	if err != nil {
		return
	}
	DeleteCourseService(ID)

	response.Deleted(c)
}

func getCourseById(c *gin.Context) {
	err, course := getCourseByIdValidation(c)
	if err != nil {
		return
	}

	response.Ok(c, course)
}
func getAllCourses(c *gin.Context) {
	err := getAllCoursesValidation(c)
	if err != nil {
		return
	}
	courses := ShowAllCoursesServices()
	response.Ok(c, courses)
}
