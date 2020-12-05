# Section 6 User api delete edit and emcrypt

## Encrypt by scrypt

**src/utils/scrypt.go:**

```go
package utils

import (
  "encoding/base64"
  "log"

  "golang.org/x/crypto/scrypt"
)

func ScryptPassword(password string) string {
  const keyLen = 10
  salt := []byte{53, 33, 26, 234, 123, 234, 42, 54}
  dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, keyLen)

  if err != nil {
    log.Fatal(err.Error())
    return password
  }
  ret := base64.StdEncoding.EncodeToString(dk)
  return ret
}
```

## Delete user Api

**src/model/user.go:**

```go

// SoftDeletUser delete user softy
func SoftDeletUser(id int) int {
  err := db.Where("id = ?", id).Delete(&User{}).Error
  if err != nil {
    return errors.ERROR
  }
  return errors.SUCCESS
}


```

**src/api/v1/user.go:**

```go
// DeleteUser delete user
func DeleteUser(c *gin.Context) {
  code := errors.SUCCESS
  id, err := strconv.Atoi(c.Query("userID"))
  if err != nil {
    code = errors.ERROR_USER_NOT_FOUND
  }
  code = model.SoftDeletUser(id)
  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   "",
    "msg":    errors.Msg(code),
  })
}
```

## Edit user Api

**src/model/user.go:**

```go
// UpdateUser update user info at username and role
func UpdateUser(id int, user *User) int {

  uMap := map[string]interface{}{
    "username": user.Username,
    "role":     user.Role,
  }

  err := db.Model(&User{}).Where("id = ?", id).Update(uMap).Error
  if err != nil {
    log.Println(err.Error())
    log.Fatal(err.Error())
    return errors.ERROR
  }
  return errors.SUCCESS
}
```

**src/api/v1/user.go:**

```go

// EditUser edit user info
func EditUser(c *gin.Context) {
  var user model.User
  id, _ := strconv.Atoi(c.Param("id"))
  c.ShouldBindJSON(&user)
  code := model.CheckUsername(user.Username)
  if code == errors.SUCCESS {
    model.UpdateUser(id, &user)
  }

  if code == errors.ERROR_USERNAME_USED {
    c.Abort()
  }

  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   "",
    "msg":    errors.Msg(code),
  })

}
```

## Search user by id

**src/api/v1/user.go:**

```go

// GetUsetByID search user by id
func GetUsetByID(c *gin.Context) {
  code := errors.SUCCESS
  msg := ""
  id, _ := strconv.Atoi(c.Param("id"))
  user, err := model.GetUsetByID(id)
  if err != nil {
    code = errors.ERROR
    msg = msg + err.Error()
  }

  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   *user,
    "msg":    msg + errors.Msg(code),
  })
}
```

**src/model/user.go:**

```go
// GetUsetByID search user with specified id
func GetUsetByID(id int) (*User, error) {
  var user User
  err := db.Where("id = ?", id).First(&user).Error
  if err != nil {
    log.Fatal(err.Error())
    return nil, err
  }
  return &user, nil
}

```
