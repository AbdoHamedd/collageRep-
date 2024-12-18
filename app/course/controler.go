package course

import (
	"basicCrudoperations/exchanges/courseExchabge/CommonPagination"
	"basicCrudoperations/response"
	"basicCrudoperations/tranform"
	"github.com/gin-gonic/gin"
)

//في الريفيرنس يا اما بتعمل ال * و متعملش ريتيرن يا اما تعمل ريتيرن و متعملش ال *

func CreateCourse(c *gin.Context) {
	err, request := CreateCourseValidation(c)
	if err != nil {
		return
	}
	course := CreateCourseService(request)
	coursetransform := tranform.CourseTransform(course)
	response.Created(c, coursetransform)
}

func updateCourse(c *gin.Context) {
	err, request, course := updateCourseValidation(c)
	if err != nil {
		return
	}
	course = UpdateCourseService(request, course)
	coursetransform := tranform.CourseTransform(course)
	response.Ok(c, coursetransform)
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
	coursetransform := tranform.CourseTransform(course)
	response.Ok(c, coursetransform)
}

func getAllCourses(c *gin.Context) {
	err, req := CommonPagination.ShawAllPaginationWithOrder(c)
	if err != nil {
		return
	}
	courses := ShowAllCoursesServices(req.Limit, req.Offset, req.Order)
	coursetransform := tranform.CoursesTransform(courses)
	response.Ok(c, coursetransform)

}
