# Section 11 Login Api

## use jwt middleware and define login api

```go
package routes

import (
  v1 "gitee.com/panda8xy/gin-blog/api/v1"
  middleware "gitee.com/panda8xy/gin-blog/middleware/jwt"
  "gitee.com/panda8xy/gin-blog/utils"
  "github.com/gin-gonic/gin"
)

// InitRouter gin engine init and run
func InitRouter() {
  gin.SetMode(utils.AppMode)

  r := gin.Default()

  authRouter := r.Group("api/v1")
  authRouter.Use(middleware.Jwt())
  {
    // user
    authRouter.POST("user/add", v1.AddUser)
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

  }

  publicRouter := r.Group("api/v1")

  {
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

```

## define login api

```go
package v1

import (
  "net/http"

  middleware "gitee.com/panda8xy/gin-blog/middleware/jwt"
  "gitee.com/panda8xy/gin-blog/model"
  "gitee.com/panda8xy/gin-blog/utils/errors"
  "github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
  var user model.User
  c.ShouldBindJSON(&user)

  var token string
  var code int

  code = model.CheckLogin(user.Username, user.Password)

  if code == errors.SUCCESS {
    token, code = middleware.GenerateToken(user.Username, user.Password)
  }
  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   gin.H{"token": token},
    "msg":    errors.Msg(code),
  })
}

```

## model func for login api

```go

// CheckLogin check user with username
func CheckLogin(username, password string) int {
  var user User
  db.Where("username = ?", username).Find(&user)
  if user.ID == 0 {
    return errors.ERROR_USER_NOT_FOUND
  }

  if utils.ScryptPassword(password) != user.Password {
    return errors.ERROR_PASSWORD_WRONG
  }

  if user.Role != 0 {
    return errors.ERROR_USER_NO_RIGHTS
  }

  return errors.SUCCESS
}
```

## add a new error code and msg

```go
ERROR_USER_NO_RIGHTS   = 1008 //  no right login backened
var msgMap = map[int]string{
    ERROR_USER_NO_RIGHTS:    "用户没有权限",
}
```
