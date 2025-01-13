package user

import (
	"basicCrudoperations/auth"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	r := router.Group("/users")
	r.POST("/signup", signup)
	r.POST("/login", login)
	r.POST("/logout", logout).Use(auth.Auth)

}
