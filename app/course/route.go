package course

import (
	"github.com/gin-gonic/gin"
)

func CourseRoutes(router *gin.Engine) {
	r := router.Group("/course")
	r.POST("/create", CreateCourse)   // POST for course creation
	r.PUT("/update", updateCourse)    // PUT for updating a course
	r.DELETE("/delete", deleteCourse) // DELETE for deleting a course
	r.GET("/getcourse", getCourseById)
	r.GET("/getallcourses", getAllCourses)

}
