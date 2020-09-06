package routes

import (
	"net/http"

	"gitee.com/panda8xy/gin-blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)

	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "OK",
			})
		})
	}

	r.Run(utils.HttpPort)

}
