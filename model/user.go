package model

import (
	"log"

	"gitee.com/panda8xy/gin-blog/utils"
	"gitee.com/panda8xy/gin-blog/utils/errors"
	"github.com/jinzhu/gorm"
)

// User user model in database and memory
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20)" json:"username"`
	Password string `gorm:"type:varchar(20)" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// CheckUsername  check username is existed
// SELECT id FROM `user`  WHERE `user`.`deleted_at` IS NULL AND ((username = 'panda')) ORDER BY `user`.`id` ASC LIMIT 1
func CheckUsername(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errors.ERROR_USERNAME_USED // 1001
	}
	return errors.SUCCESS
}

// GetUsetByID search user with specified id
func GetUsetByID(id int) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &user, nil
}

// CreateUser add a new user to database
// INSERT INTO `user` (`created_at`,`updated_at`,`deleted_at`,`username`,`password`,`role`) VALUES ('2020-09-07 12:40:05','2020-09-07 12:40:05',NULL,'panda','123456',0)
func CreateUser(user *User) int {
	// user.Password = ScryptPassword(user.Password)
	err := db.Create(&user).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

// GetUserList get user list by pageable
// SELECT * FROM `user`  WHERE `user`.`deleted_at` IS NULL LIMIT 20 OFFSET 0
func GetUserList(pageSize int, pageNum int) []User {
	var users []User

	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// SoftDeletUser delete user softy
func SoftDeletUser(id int) int {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

// UpdateUser update user info at username and role
func UpdateUser(id int, user *User) int {

	uMap := map[string]interface{}{
		"username": user.Username,
		"role":     user.Role,
	}

	err := db.Model(&User{}).Where("id = ?", id).Update(uMap).Error
	if err != nil {
		log.Println(err.Error())
		log.Fatal(err.Error())
		return errors.ERROR
	}
	return errors.SUCCESS
}

// BeforeSave ecrypt use scrypt
func (u *User) BeforeSave() {
	u.Password = utils.ScryptPassword(u.Password)
}
