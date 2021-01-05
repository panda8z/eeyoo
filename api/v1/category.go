package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/eeyoo/model"
	"github.com/panda8z/eeyoo/utils/errors"
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
	data, total := model.CategoryList(pageSize, pageNum)
	code := errors.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": gin.H{
			"list":  data,
			"total": total,
		},
		"msg": errors.Msg(code),
	})

}

// GetCateByID search cate by id··
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
