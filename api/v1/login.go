package v1

import (
	"net/http"

	"gitee.com/panda8xy/gin-blog/middleware"
	"gitee.com/panda8xy/gin-blog/model"
	"gitee.com/panda8xy/gin-blog/utils/errors"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	var token string
	var code int

	code = model.CheckLogin(user.Username, user.Password)

	if code == errors.SUCCESS {
		token, code = middleware.GenerateToken(user.Username, user.Password)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   gin.H{"token": token},
		"msg":    errors.Msg(code),
	})
}
