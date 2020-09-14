package v1

import (
	"net/http"

	"gitee.com/panda8xy/gin-blog/model"
	"gitee.com/panda8xy/gin-blog/utils/errors"
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
