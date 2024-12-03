package routes

import (
	"basicCrudoperations/app/course"
	"basicCrudoperations/app/department"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	course.CourseRoutes(r)
	department.DepartmentRoute(r)

}
