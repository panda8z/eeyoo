package model

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/panda8z/eeyoo/utils/errors"
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
func CategoryList(pageSize int, pageNum int) ([]Category, int) {
	var cates []Category
	var total int

	db.Model(&cates).Count(&total)
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cates, total
}

// SoftDeletCategory delete cate softy
func SoftDeletCategory(id int) int {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

// UpdateCategoryName update cate info at name
func UpdateCategoryName(id int, cate *Category) int {
	uMap := map[string]interface{}{
		"name": cate.Name,
	}
	// UPDATE `category` SET `name` = 'aaa1'  WHERE (id = 0)
	err := db.Model(&cate).Where("id = ?", id).Updates(uMap).Error
	if err != nil {
		log.Println(err.Error())
		log.Fatal(err.Error())
		return errors.ERROR
	}
	return errors.SUCCESS
}
