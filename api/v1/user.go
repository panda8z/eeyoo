package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/eeyoo/model"
	"github.com/panda8z/eeyoo/utils/errors"
	"github.com/panda8z/eeyoo/utils/validate"
)

// AddUser add user
func AddUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	msg, code := validate.Validate(&user)
	if code != errors.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    msg,
		})
		return
	}
	code = model.CheckUsername(user.Username)
	if code == errors.SUCCESS {
		model.CreateUser(&user)
	}

	if code == errors.ERROR_USERNAME_USED {
		code = errors.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errors.Msg(code),
	})
}

// GetUserList search user list
func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	username := c.Query("username")
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total := model.GetUserList(username, pageSize, pageNum)
	code := errors.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": gin.H{
			"list":  data,
			"total": total,
		},
		"msg": errors.Msg(code),
	})
}

// GetUsetByID search user by id
func GetUsetByID(c *gin.Context) {
	code := errors.SUCCESS
	msg := ""
	id, _ := strconv.Atoi(c.Param("id"))
	err, user := model.GetUsetByID(id)
	if err != nil {
		code = errors.ERROR
		msg = err.Error()
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   *user,
		"msg":    msg + errors.Msg(code),
	})
}

// EditUser edit user info
func EditUser(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&user)

	code := model.CheckUpUser(id, user.Username)

	if code == errors.ERROR_USERNAME_USED {
		//c.Abort()
	}

	if code == errors.SUCCESS {
		model.UpdateUser(id, &user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errors.Msg(code),
	})
}

// DeleteUser delete user
func DeleteUser(c *gin.Context) {
	code := errors.SUCCESS
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.SoftDeletUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errors.Msg(code),
	})
}
