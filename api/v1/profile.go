package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/eeyoo/model"
	"github.com/panda8z/eeyoo/utils/errors"
)

func GetProfile(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	profile, code := model.GetProfile(id)
	c.JSON(http.StatusOK, gin.H{
		"data":   profile,
		"status": code,
		"msg":    errors.Msg(code),
	})

}

func EditProfile(c *gin.Context) {
	var profile model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&profile)
	code := model.UpdateProfile(id, &profile)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errors.Msg(code),
	})
}
