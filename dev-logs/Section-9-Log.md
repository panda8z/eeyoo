
# Section 9 Article list and info api

## article api

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

// GetArticleInfo search article by id
func GetArticleInfo(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  article, code := model.ArticleInfoByID(id)
  if code == errors.ERROR_ARTICLE_NOT_EXIST {
    c.Abort()
  }
  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   article,
    "msg":    errors.Msg(code),
  })
}

// GetArticlesByCateID search article by id
func GetArticlesByCateID(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  pageSize, _ := strconv.Atoi(c.Query("pageSize"))
  pageNum, _ := strconv.Atoi(c.Query("pageNum"))

  if pageSize == 0 {
    pageSize = -1
  }
  if pageNum == 0 {
    pageNum = -1
  }
  art, code := model.ArticlesByCateID(id, pageSize, pageNum)
  if code == errors.ERROR_ARTICLE_NOT_EXIST {
    c.Abort()
  }
  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   *art,
    "msg":    errors.Msg(code),
  })
}

// GetArticleList search article list in pageable
func GetArticleList(c *gin.Context) {
  pageSize, _ := strconv.Atoi(c.Query("pageSize"))
  pageNum, _ := strconv.Atoi(c.Query("pageNum"))

  if pageSize == 0 {
    pageSize = -1
  }
  if pageNum == 0 {
    pageNum = -1
  }
  data, code := model.ArticleList(pageSize, pageNum)
  c.JSON(http.StatusOK, gin.H{
    "status": code,
    "data":   data,
    "msg":    errors.Msg(code),
  })

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

## article model

```go
package model

import (
  "log"

  "github.com/panda8z/eeyoo/utils/errors"
  "github.com/jinzhu/gorm"
)

// Article model
type Article struct {
  Category Category `gorm:"foreignkey:Cid"`
  gorm.Model
  Title   string `gorm:"type:varchar(100);not null" json:"title"`
  Cid     int    `gorm:"type:int;not null" json:"cid"`
  Desc    string `gorm:"type:varchar(200)" json:"desc"`
  Content string `gorm:"type:longtext" json:"content"`
  Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArticle Add a new art to database
func CreateArticle(art *Article) int {
  err := db.Create(&art).Error
  if err != nil {
    return errors.ERROR
  }
  return errors.SUCCESS
}

// CheckArticleName  check article name is existed
// func CheckArticleName(name string) int {
//   var art Article
//   db.Select("id").Where("name = ?", name).First(&art)
//   if art.ID > 0 {
//     return errors.ERROR_CATEGORY_USED // 2001
//   }
//   return errors.SUCCESS
// }

// ArticlesByCateID getter for articles with same Category ID
func ArticlesByCateID(id int, pageSize int, pageNum int) (*[]Article, int) {
  var arts []Article
  err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).First(&arts).Error
  if err != nil {
    log.Fatal(err.Error())
    return nil, errors.ERROR_ARTICLE_NOT_EXIST
  }
  return &arts, errors.SUCCESS
}

// ArticleInfoByID getter for article with specified Article ID
func ArticleInfoByID(id int) (Article, int) {
  var art Article
  err := db.Preload("Category").Where("id = ?", id).First(&art).Error
  if err != nil {
    log.Fatal(err.Error())
    return art, errors.ERROR_ARTICLE_NOT_EXIST
  }
  return art, errors.SUCCESS
}

// ArticleList get art list in pageable
func ArticleList(pageSize int, pageNum int) ([]Article, int) {
  // SELECT * FROM `article`  WHERE `article`.`deleted_at` IS NULL LIMIT 20 OFFSET 0
  // SELECT * FROM `category`  WHERE `category`.`deleted_at` IS NULL AND ((`id` IN (5)))
  var arts []Article
  err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
  if err != nil && err != gorm.ErrRecordNotFound {
    return nil, errors.ERROR
  }
  return arts, errors.SUCCESS
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

## 参考资料

- [mysql中如何在创建数据库的时候指定数据库的字符集? - QA-3K - 博客园](https://www.cnblogs.com/chuanzhang053/p/9121506.html)
