package user

import "github.com/gin-gonic/gin"

func userRoute(route *gin.Engine) {
	r := route.Group("/user")
	r.POST("/signUp", signUp)
}
