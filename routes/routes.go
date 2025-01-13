package routes

import (
	"basicCrudoperations/app/course"
	"basicCrudoperations/app/department"
	"basicCrudoperations/app/student_course"
	"basicCrudoperations/app/user"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	course.CourseRoutes(r)
	department.DepartmentRoute(r)
	user.UserRoute(r)
	student_course.StudentCourseRoute(r)

}
