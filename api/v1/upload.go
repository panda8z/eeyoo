package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/eeyoo/model"
	"github.com/panda8z/eeyoo/utils/errors"
)

func UploadFile(c *gin.Context) {
	file, fHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.Abort()
		return
	}
	fileSize := fHeader.Size

	url, code := model.UploadFile(file, fileSize)
	url = strings.Replace(url, "qgmpmk7gb", "qmr6raydr", -1)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": gin.H{
			"url": url,
		},
		"msg": errors.Msg(code),
	})
}
