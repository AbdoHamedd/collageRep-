package auth

import (
	"basicCrudoperations/response"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Auth(c *gin.Context) {
	// todo get the Token
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	count := 0
	for _, i2 := range tokenString {
		if i2 == '.' {
			count++
		}
	}
	if count != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//todo Decode & Validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("1a2b3d4o5h6a7m8e6d"), nil
	})

	// todo Check The Exp
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > float64(claims["exp"].(float64)) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//todo Continue
		c.Next()
	}
	c.AbortWithStatus(http.StatusUnauthorized)

}
func AuthAdmin(c *gin.Context) {
	// todo get the role header
	role := c.GetHeader("role")
	if role != "Admin" {
		c.AbortWithStatus(http.StatusUnauthorized)
		response.BadRequest(c, "you can't access this")
	}
	c.Next()
}
