package model

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/panda8z/eeyoo/utils/errors"
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
// 	var art Article
// 	db.Select("id").Where("name = ?", name).First(&art)
// 	if art.ID > 0 {
// 		return errors.ERROR_CATEGORY_USED // 2001
// 	}
// 	return errors.SUCCESS
// }

// ArticlesByCateID getter for articles with same Category ID
func ArticlesByCateID(id int, pageSize int, pageNum int) (*[]Article, int32, int) {
	var arts []Article
	var total int32
	db.Preload("Category").Model(&arts).Where("cid = ?", id).Count(&total)
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&arts).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, 0, errors.ERROR_CATEGORY_USED
	}
	return &arts, total, errors.SUCCESS
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
func ArticleList(title string, pageSize int, pageNum int) ([]Article, int, int) {
	var arts []Article
	var total int

	if title != "" {
		db.Model(&arts).Where("title LIKE ?", title+"%").Count(&total)
		err := db.Order("Updated_At Desc").Preload("Category").Where("title LIKE ?", title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0, errors.ERROR
		}
	} else {
		db.Model(&arts).Count(&total)
		err := db.Order("Updated_At Desc").Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0, errors.ERROR
		}
	}

	return arts, total, errors.SUCCESS
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
