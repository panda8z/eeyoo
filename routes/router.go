package routes

import (
	"net/http"

	v1 "gitee.com/panda8xy/gin-blog/api/v1"
	"gitee.com/panda8xy/gin-blog/utils"
	"github.com/gin-gonic/gin"
)

// InitRouter gin engine init and run
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

		// user
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUsetByID)
		router.GET("users", v1.GetUserList)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		//category
		router.POST("cate/add", v1.AddCate)
		router.GET("cates", v1.GetCategoryList)
		router.GET("cate/:id", v1.GetCateByID)
		router.PUT("cate/:id", v1.EditCate)
		router.DELETE("cate/:id", v1.DeleteCate)

		// article

		// lgoin
	}

	r.Run(utils.HttpPort)

}
