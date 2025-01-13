package student_course

import (
	"basicCrudoperations/exchanges/studenr_course"
	"basicCrudoperations/response"
	"basicCrudoperations/tranform"
	"github.com/gin-gonic/gin"
)

func StudentJoinCourseControl(c *gin.Context) {
	var req studenr_course.StudentJoinCourseRequest
	check := StudentJoinCourseValidate(c)
	if !check {
		return
	}
	studentJoinCourse := studentJoinCourseService(req)
	res := tranform.TransformJoin(&studentJoinCourse)
	response.Ok(c, res)

}

func StudendDrobCourseController(c *gin.Context) {
	var req studenr_course.StudentJoinCourseRequest
	check := StudendDrobCourseValidate(c)
	if !check {
		return
	}
	JoinDrobService(req)
	response.Deleted(c)

}
