package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginInfo struct {
	Mail     string `validate:"required,email"`
	Password string
}

var validate *validator.Validate

// 登录
func Login(c *gin.Context) {
	// params
	mail := c.PostForm("mail")
	password := c.PostForm("password")

	// validate
	validate = validator.New()
	errs := validate.Var(mail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "请输入正确的邮箱！",
		})
		return
	}

	var loginInfo LoginInfo

	loginInfo.Mail = mail
	loginInfo.Password = password

	c.SecureJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": &loginInfo,
	})
}
