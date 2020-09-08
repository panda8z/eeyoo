package v1

import (
	"net/http"
	"strconv"

	"gitee.com/panda8xy/gin-blog/model"
	"gitee.com/panda8xy/gin-blog/utils/errors"
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
