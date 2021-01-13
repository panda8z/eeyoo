package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/eeyoo/model"
	"github.com/panda8z/eeyoo/utils/errors"
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
	art, total, code := model.ArticlesByCateID(id, pageSize, pageNum)

	if code != errors.SUCCESS {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": gin.H{
			"list":  art,
			"total": total,
		},
		"msg": errors.Msg(code),
	})
}

// GetArticleList search article list in pageable
func GetArticleList(c *gin.Context) {
	title := c.Query("title")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, total, code := model.ArticleList(title, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": gin.H{
			"list":  data,
			"total": total,
		},
		"msg": errors.Msg(code),
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
