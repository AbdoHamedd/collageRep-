package student_course

import (
	"basicCrudoperations/auth"
	"github.com/gin-gonic/gin"
)

func StudentCourseRoute(router *gin.Engine) {
	r := router.Group("student_course").Use(auth.Auth)
	r.POST("/join", StudentJoinCourseControl)
	r.DELETE("/drob", StudendDrobCourseController)
}
