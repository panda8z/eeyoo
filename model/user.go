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
	Username string `gorm:"type:varchar(20)" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20)" json:"password" validate:"required,min=6,max=22" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"`
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
func GetUserList(pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
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

// CheckLogin check user with username
func CheckLogin(username, password string) int {
	var user User
	db.Where("username = ?", username).Find(&user)
	if user.ID == 0 {
		return errors.ERROR_USER_NOT_FOUND
	}

	if utils.ScryptPassword(password) != user.Password {
		return errors.ERROR_PASSWORD_WRONG
	}

	if user.Role != 1 {
		return errors.ERROR_USER_NO_RIGHTS
	}

	return errors.SUCCESS
}
