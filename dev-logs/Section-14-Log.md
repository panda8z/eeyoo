# Section 14 Validation and Modify User-Add func

- github.com/go-playground/validator/v10

## add user struct `validate` tag

[validator package · pkg.go.dev](https://pkg.go.dev/github.com/go-playground/validator/v10?tab=doc)

```go
// User user model in database and memory
type User struct {
  gorm.Model
  Username string `gorm:"type:varchar(20)" json:"username" validate:"required,min=4,max=12"`
  Password string `gorm:"type:varchar(20)" json:"password" validate:"required,min=6,max=22"`
  Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2"`
}
```

## change the login model func

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

  if user.Role != 1 {
    return errors.ERROR_USER_NO_RIGHTS
  }

  return errors.SUCCESS
}

```

## validate

```bash
go get -u github.com/go-playground/validator/v10
go get -u github.com/go-playground/universal-translator
go get -u github.com/go-playground/locales
```

```go
package validate

import (
  "fmt"

  "github.com/panda8z/eeyoo/utils/errors"
  "github.com/go-playground/locales/zh_Hans_CN"
  trans "github.com/go-playground/universal-translator"
  "github.com/go-playground/validator/v10"
  zh "github.com/go-playground/validator/v10/translations/zh"
)

func Validate(data interface{}) (string, int) {
  v := validator.New()
  u := trans.New(zh_Hans_CN.New())
  t, _ := u.GetTranslator(zh_Hans_CN.New().Locale())
  err := zh.RegisterDefaultTranslations(v, t)
  if err != nil {
    fmt.Println(err.Error())
  }
  err = v.Struct(data)
  if err != nil {
    for _, val := range err.(validator.ValidationErrors) {
      return val.Translate(t), errors.ERROR
    }
  }
  return "", errors.SUCCESS
}

```

## change the add user api

```go
// AddUser add user
func AddUser(c *gin.Context) {
  var user model.User
  _ = c.ShouldBindJSON(&user)
  msg, code := validate.Validate(&user)
  if code != errors.SUCCESS {
    c.JSON(http.StatusOK, gin.H{
      "status": code,
      "msg":    msg,
    })
    return
  }
  code = model.CheckUsername(user.Username)
  if code == errors.SUCCESS {
    model.CreateUser(&user)
  }

  if code == errors.ERROR_USERNAME_USED {
    code = errors.ERROR_USERNAME_USED
  }

  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "msg":    errors.Msg(code),
  })
}
```

## add user label

```go
// User user model in database and memory
type User struct {
  gorm.Model
  Username string `gorm:"type:varchar(20)" json:"username" validate:"required,min=4,max=12" label:"用户名"`
  Password string `gorm:"type:varchar(20)" json:"password" validate:"required,min=6,max=22" label:"密码"`
  Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"`
}

```

```go

```
