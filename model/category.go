package model

import (
	"log"

	"gitee.com/panda8xy/gin-blog/utils/errors"
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
