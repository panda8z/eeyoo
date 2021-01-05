package v1

import (
	"net/http"

	"github.com/panda8z/eeyoo/middleware"
	"github.com/panda8z/eeyoo/model"
	"github.com/panda8z/eeyoo/utils/errors"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)

	var token string
	var code int

	code = model.CheckLogin(user.Username, user.Password)

	if code == errors.SUCCESS {
		token, code = middleware.GenerateToken(user.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   gin.H{"token": token},
		"msg":    errors.Msg(code),
	})
}
