package user

//
//import (
//	"basicCrudoperations/exchanges"
//	"basicCrudoperations/response"
//	"errors"
//	"github.com/gin-gonic/gin"
//	"github.com/thedevsaddam/govalidator"
//	"net/http"
//	"net/url"
//)
//
//func signUpValidate(c *gin.Context) bool {
//	var req exchanges.SignUpRequest
//	err := c.ShouldBind(&req)
//	if err != nil {
//		response.BadRequest(c, err.Error())
//		return false
//	}
//	check, err := ValidateInputsForSignUp(c.Request)
//	if !check {
//		err = errors.New("Inputs Is Invalid")
//		response.BadRequest(c, err.Error())
//		return false
//	}
//	return true
//
//}
//func ValidateInputsForSignUp(c *http.Request) (bool, url.Values) {
//	var req exchanges.SignUpRequest
//	rules := govalidator.MapData{
//		"email":    []string{"required", "email", "unique:users"},
//		"Password": []string{"required"},
//		"username": []string{"required"},
//	}
//	options := govalidator.Options{
//		Request:         c,
//		Rules:           rules,
//		Data:            req,
//		RequiredDefault: true,
//	}
//	createNewVal := govalidator.New(options)
//	err := createNewVal.ValidateJSON()
//	if err != nil {
//		return false, err
//
//	}
//	return true, err
//
//}
