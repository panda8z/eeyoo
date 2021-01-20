package model

import (
	"github.com/jinzhu/gorm"
	"github.com/panda8z/eeyoo/utils/errors"
)

type Profile struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20)" json:"name"`
	Desc      string `gorm:"type:varchar(200)" json:"desc"`
	QQchat    string `gorm:"type:varchar(20)" json:"qqchat"`
	Wechat    string `gorm:"type:varchar(20)" json:"wechat"`
	Weibo     string `gorm:"type:varchar(40)" json:"weibo"`
	Bili      string `gorm:"type:varchar(20)" json:"bili"`
	Email     string `gorm:"type:varchar(20)" json:"email"`
	Img       string `gorm:"type:varchar(20)" json:"img"`
	Avatar    string `gorm:"type:varchar(20)" json:"avatar"`
	Facebook  string `gorm:"type:varchar(20)" json:"facebook"`
	Twitter   string `gorm:"type:varchar(20)" json:"twitter"`
	Instagram string `gorm:"type:varchar(20)" json:"instagram"`
}

// 获取个人信息
func GetProfile(id int) (Profile, int) {
	var profile Profile
	err = db.Where("ID = ?", id).First(&profile).Error
	if err != nil {
		return profile, errors.ERROR
	}
	return profile, errors.SUCCESS
}

// 整体更新个人信息
// TODO：改成拆开更新，并继续更新的具体字段
func UpdateProfile(id int, data *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}
