package routes

import (
	"time"

	v1 "gitee.com/panda8xy/gin-blog/api/v1"
	"gitee.com/panda8xy/gin-blog/middleware"
	"gitee.com/panda8xy/gin-blog/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRouter gin engine init and run
func InitRouter() {
	gin.SetMode(utils.AppMode)

	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"https://foo.com"},
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		// AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))

	authRouter := r.Group("api/v1")
	authRouter.Use(middleware.Jwt())
	{
		// user
		authRouter.PUT("user/:id", v1.EditUser)
		authRouter.DELETE("user/:id", v1.DeleteUser)

		//category
		authRouter.POST("cate/add", v1.AddCate)
		authRouter.PUT("cate/:id", v1.EditCate)
		authRouter.DELETE("cate/:id", v1.DeleteCate)

		// article
		authRouter.POST("article/add", v1.AddArticle)
		authRouter.PUT("article/:id", v1.EditArt)
		authRouter.DELETE("article/:id", v1.DeleteArt)

		// upload file
		authRouter.POST("upload", v1.UploadFile)

	}

	publicRouter := r.Group("api/v1")

	{
		// 注册?
		publicRouter.POST("user/add", v1.AddUser)
		// lgoin
		publicRouter.POST("login", v1.Login)
		// user
		publicRouter.GET("user/:id", v1.GetUsetByID)
		publicRouter.GET("users", v1.GetUserList)

		//category
		publicRouter.GET("cates", v1.GetCategoryList)
		publicRouter.GET("cate/:id", v1.GetCateByID)

		// article
		publicRouter.GET("article/list", v1.GetArticleList)
		publicRouter.GET("article/cate/:id", v1.GetArticlesByCateID)
		publicRouter.GET("article/info/:id", v1.GetArticleInfo)
	}

	r.Run(utils.HttpPort)

}
