# Section 3 V1/Api structrue and  Errors map

## Error Map

```go
package errors

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000 - 1999 user errs
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_FOUND   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_OUTTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007

	// code = 2000 - 2999 article errs

	// code = 3000 - 3999 category errs
)

var codeMap = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FIAL",
	ERROR_USERNAME_USED:    "用户名已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_FOUND:   "用户未找到",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_OUTTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
}

// Code getter for errors
func Code(codeNum int) string {
	return codeMap[codeNum]
}

```

## v1/apis

` cd api/v1 && touch article.go category.go login.go user.go `

**src/api/v1/user.go:**

```go
package v1

import "github.com/gin-gonic/gin"

// search user existed
func UserExisted(c *gin.Context) {

}

// add user
func AddUser(c *gin.Context) {

}

// search user by id
func GetUsetById(c *gin.Context) {

}

// search user list
func GetUserList(c *gin.Context) {

}

// edit user info
func EditUser(c *gin.Context) {

}

// delete user
func DeleteUser(c *gin.Context) {

}
```


## router structure

```go 
package routes

import (
	"net/http"

	v1 "github.com/panda8z/eeyoo/api/v1"
	"github.com/panda8z/eeyoo/utils"
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

		// user
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUsetById)
		router.GET("users", v1.GetUserList)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)

		// article

		// lgoin
	}

	r.Run(utils.HttpPort)

}

```