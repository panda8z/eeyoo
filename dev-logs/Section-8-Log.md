# Section 8 Article add edit delet api

## api

```go
package v1

import (
  "net/http"
  "strconv"

  "github.com/panda8z/eeyoo/model"
  "github.com/panda8z/eeyoo/utils/errors"
  "github.com/gin-gonic/gin"
)

// AddArticle article add
func AddArticle(c *gin.Context) {
  var art model.Article
  _ = c.ShouldBindJSON(&art)
  code := model.CreateArticle(&art)

  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   art,
    "msg":    errors.Msg(code),
  })
}

// GetArticleByID search article by id
func GetArticleByID(c *gin.Context) {

}

// GetArticleList search article list in pageable
func GetArticleList(c *gin.Context) {
}

// EditArt edit art info
func EditArt(c *gin.Context) {
  var art model.Article
  id, _ := strconv.Atoi(c.Param("id"))
  c.ShouldBindJSON(&art)
  code := model.UpdateArticle(id, &art)
  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   "",
    "msg":    errors.Msg(code),
  })

}

// DeleteArt delete art
func DeleteArt(c *gin.Context) {
  code := errors.SUCCESS
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    code = errors.ERROR_USER_NOT_FOUND
  }
  code = model.SoftDeletArticle(id)
  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   "",
    "msg":    errors.Msg(code),
  })
}

```

## model methods

```go
package model

import (
  "log"

  "github.com/panda8z/eeyoo/utils/errors"
  "github.com/jinzhu/gorm"
)

// Article model
type Article struct {
  gorm.Model
  Title   string `gorm:"type:varchar(100);not null" json:"title"`
  Cid     int    `gorm:"type:int;not null" json:"cid"`
  Desc    string `gorm:"type:varchar(200)" json:"desc"`
  Content string `gorm:"type:longtext" json:"content"`
  Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CheckArticleName  check name is existed
func CheckArticleName(name string) int {
  var art Article
  db.Select("id").Where("name = ?", name).First(&art)
  if art.ID > 0 {
    return errors.ERROR_CATEGORY_USED // 2001
  }
  return errors.SUCCESS
}

// GetArticleByID search art with specified id
func GetArticleByID(id int) (*Article, error) {
  var art Article
  err := db.Where("id = ?", id).First(&art).Error
  if err != nil {
    log.Fatal(err.Error())
    return nil, err
  }
  return &art, nil
}

// CreateArticle add a new art to database
func CreateArticle(art *Article) int {
  err := db.Create(&art).Error
  if err != nil {
    return errors.ERROR
  }
  return errors.SUCCESS
}

// ArticleList get art list in pageable
func ArticleList(pageSize int, pageNum int) []Article {
  var arts []Article

  err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return nil
  }
  return arts
}

// SoftDeletArticle delete art softy
func SoftDeletArticle(id int) int {
  err := db.Where("id = ?", id).Delete(&Article{}).Error
  if err != nil {
    return errors.ERROR
  }
  return errors.SUCCESS
}

// UpdateArticle update art info at name and role
func UpdateArticle(id int, art *Article) int {
  uMap := map[string]interface{}{
    "Title":   art.Title,
    "Desc":    art.Desc,
    "Content": art.Content,
  }

  code := CheckCateName(art.Title)
  if code == errors.ERROR_CATEGORY_USED {
    return errors.ERROR_CATEGORY_USED
  }

  err := db.Model(&Article{}).Where("id = ?", id).Update(uMap).Error
  if err != nil {
    log.Println(err.Error())
    log.Fatal(err.Error())
    return errors.ERROR
  }
  return errors.SUCCESS
}

```

## add in router

```go
// article
router.POST("article/add", v1.AddArticle)    //
router.GET("articles", v1.GetArticleList)    // todo
router.GET("article/:id", v1.GetArticleByID) // todo
router.PUT("article/:id", v1.EditArt)        //
router.DELETE("article/:id", v1.DeleteArt)   //
```
