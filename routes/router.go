package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/panda8z/eeyoo/api/v1"
	"github.com/panda8z/eeyoo/middleware"
	"github.com/panda8z/eeyoo/utils"
)

// InitRouter gin engine init and run
func InitRouter() {
	gin.SetMode(utils.AppMode)

	r := gin.New()

	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CorsConfig())
	r.LoadHTMLGlob("web/admin/dist/index.html")
	r.Static("admin/static", "web/admin/dist/static")
	r.StaticFile("admin/favicon.ico", "web/admin/dist/favicon.ico")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

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

		authRouter.POST("profile/:id", v1.EditProfile)

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

		authRouter.GET("profile/:id", v1.GetProfile)
	}

	r.Run(utils.HttpPort)

}
