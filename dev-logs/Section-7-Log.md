# Section 7 Category Api and model methods

## Cates Api

```go
package v1

import (
  "net/http"
  "strconv"

  "github.com/panda8z/eeyoo/model"
  "github.com/panda8z/eeyoo/utils/errors"
  "github.com/gin-gonic/gin"
)

// AddCate add cate
func AddCate(c *gin.Context) {
  var cate model.Category
  _ = c.ShouldBindJSON(&cate)
  code := model.CheckCateName(cate.Name)
  if code == errors.SUCCESS {
    model.CreateCate(&cate)
  }

  if code == errors.ERROR_CATEGORY_USED {
    code = errors.ERROR_CATEGORY_USED
    cate = model.Category{}
  }

  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   cate,
    "msg":    errors.Msg(code),
  })
}

// GetCategoryList search cate list
func GetCategoryList(c *gin.Context) {
  pageSize, _ := strconv.Atoi(c.Query("pageSize"))
  pageNum, _ := strconv.Atoi(c.Query("pageNum"))

  if pageSize == 0 {
    pageSize = -1
  }
  if pageNum == 0 {
    pageNum = -1
  }
  data := model.CategoryList(pageSize, pageNum)
  code := errors.SUCCESS
  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   data,
    "msg":    errors.Msg(code),
  })

}

// GetCateByID search cate by id
func GetCateByID(c *gin.Context) {
  code := errors.SUCCESS
  msg := ""
  id, _ := strconv.Atoi(c.Param("id"))
  cate, err := model.GetCateByID(id)
  if err != nil {
    code = errors.ERROR
    msg = msg + err.Error()
  }

  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   *cate,
    "msg":    msg + errors.Msg(code),
  })
}

// EditCate edit cate info
func EditCate(c *gin.Context) {
  var cate model.Category
  id, _ := strconv.Atoi(c.Query("id"))
  c.ShouldBindJSON(&cate)
  code := model.CheckCateName(cate.Name)
  if code == errors.SUCCESS {
    model.UpdateCategory(id, &cate)
  }

  if code == errors.ERROR_CATEGORY_USED {
    c.Abort()
  }

  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   "",
    "msg":    errors.Msg(code),
  })

}

// DeleteCate delete cate
func DeleteCate(c *gin.Context) {
  code := errors.SUCCESS
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    code = errors.ERROR_USER_NOT_FOUND
  }
  code = model.SoftDeletCategory(id)
  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   "",
    "msg":    errors.Msg(code),
  })
}

```

## cates model methods

```go
package model

import (
  "log"

  "github.com/panda8z/eeyoo/utils/errors"
  "github.com/jinzhu/gorm"
)

// Category model
type Category struct {
  gorm.Model
  Name string `gorm:"type:varchar(20);not null;" json:"name"`
}

// CheckCateName  check name is existed
func CheckCateName(name string) int {
  var cate Category
  db.Select("id").Where("name = ?", name).First(&cate)
  if cate.ID > 0 {
    return errors.ERROR_CATEGORY_USED // 2001
  }
  return errors.SUCCESS
}

// GetCateByID search cate with specified id
func GetCateByID(id int) (*Category, error) {
  var cate Category
  err := db.Where("id = ?", id).First(&cate).Error
  if err != nil {
    log.Fatal(err.Error())
    return nil, err
  }
  return &cate, nil
}

// CreateCate add a new cate to database
func CreateCate(cate *Category) int {
  err := db.Create(&cate).Error
  if err != nil {
    return errors.ERROR
  }
  return errors.SUCCESS
}

// CategoryList get cate list in pageable
func CategoryList(pageSize int, pageNum int) []Category {
  var cates []Category

  err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return nil
  }
  return cates
}

// SoftDeletCategory delete cate softy
func SoftDeletCategory(id int) int {
  err := db.Where("id = ?", id).Delete(&Category{}).Error
  if err != nil {
    return errors.ERROR
  }
  return errors.SUCCESS
}

// UpdateCategory update cate info at name and role
func UpdateCategory(id int, cate *Category) int {

  uMap := map[string]interface{}{
    "name": cate.Name,
  }

  code := CheckCateName(cate.Name)
  if code == errors.ERROR_CATEGORY_USED {
    return errors.ERROR_CATEGORY_USED
  }

  err := db.Model(&Category{}).Where("id = ?", id).Update(uMap).Error
  if err != nil {
    log.Println(err.Error())
    log.Fatal(err.Error())
    return errors.ERROR
  }
  return errors.SUCCESS
}

```

## add to router.go

```go
//category
    router.POST("cate/add", v1.AddCate)
    router.GET("cates", v1.GetCategoryList)
    router.GET("cate/:id", v1.GetCateByID)
    router.PUT("cate/:id", v1.EditCate)
    router.DELETE("cate/:id", v1.DeleteCate)
```
