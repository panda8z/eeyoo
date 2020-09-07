package v1

import (
	"log"
	"net/http"
	"strconv"

	"gitee.com/panda8xy/gin-blog/model"
	"gitee.com/panda8xy/gin-blog/utils/errors"
	"github.com/gin-gonic/gin"
)

// UserExisted search user existed
func UserExisted(c *gin.Context) {

}

// AddUser add user
func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	log.Println(user)
	code := model.CheckUsername(user.Username)
	if code == errors.SUCCESS {
		model.CreateUser(&user)
	}

	if code == errors.ERROR_USERNAME_USED {
		code = errors.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   user,
		"msg":    errors.Msg(code),
	})
}

// GetUserList search user list
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data := model.GetUserList(pageSize, pageNum)
	code := errors.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errors.Msg(code),
	})

}

// GetUsetById search user by id
func GetUsetById(c *gin.Context) {

}

// edit user info
func EditUser(c *gin.Context) {

}

// delete user
func DeleteUser(c *gin.Context) {

}
