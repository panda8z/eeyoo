# Section 4 User Api model

## Check user isexisted by username

**model:**

```go
// CheckUsername  check username is existed
// SELECT id FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((username = 'panda')) ORDER BY `user`.`id` ASC LIMIT 1
func CheckUsername(name string) int {
 var user User
 db.Select("id").Where("username = ?", name).First(&user)
 if user.ID > 0 {
  return errors.ERROR_USERNAME_USED // 1001
 }
 return errors.SUCCESS
}
```

## CreateUser

**model:**

```go
// CreateUser add a new user to database
// INSERT INTO `user` (`created_at`,`updated_at`,`deleted_at`,`username`,`password`,`role`) VALUES ('2020-09-07 12:40:05','2020-09-07 12:40:05',NULL,'panda','123456',0)
func CreateUser(user *User) int {
 user.Password = scryptPassword(user.Password)
 err := db.Create(&user).Error
 if err != nil {
  return errors.ERROR
 }
 return errors.SUCCESS
}
```

**api:**

```go
// AddUser add user
func AddUser(c *gin.Context) {
 var user model.User
 _ = c.ShouldBindJSON(&user)
 log.Println(user)
 code := model.CheckUsername(user.Username)
 if code == errors.SUCCESS {
  model.CreateUser(&user)
 }

 if code == errors.ERROR_USERNAME_USED {
  code = errors.ERROR_USERNAME_USED
 }

 c.JSON(http.StatusOK, gin.H{
  "status": code,
  "data":   user,
  "msg":    errors.Msg(code),
 })
}
```

## Get user List by pageable

**model:**

```go
// GetUserList get user list by pageable
// SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL LIMIT 20 OFFSET 0
func GetUserList(pageSize int, pageNum int) []User {
 var users []User

 err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
 if err != nil && err != gorm.ErrRecordNotFound {
  return nil
 }
 return users
}

```

**api:**

```go

// GetUserList search user list
func GetUserList(c *gin.Context) {
 pageSize, _ := strconv.Atoi(c.Query("pageSize"))
 pageNum, _ := strconv.Atoi(c.Query("pageNum"))

 if pageSize == 0 {
  pageSize = -1
 }
 if pageNum == 0 {
  pageNum = -1
 }
 data := model.GetUserList(pageSize, pageNum)
 code := errors.SUCCESS
 c.JSON(http.StatusOK, gin.H{
  "status": code,
  "data":   data,
  "msg":    errors.Msg(code),
 })

}
```
