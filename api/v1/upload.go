package v1

import (
	"net/http"

	"github.com/panda8z/eeyoo/model"
	"github.com/panda8z/eeyoo/utils/errors"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, fHeader, _ := c.Request.FormFile("file")
	fileSize := fHeader.Size

	url, code := model.UploadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": gin.H{
			"url": url,
		},
		"msg": errors.Msg(code),
	})
}
