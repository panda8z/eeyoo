package model

import (
	"encoding/base64"
	"log"

	"gitee.com/panda8xy/gin-blog/utils/errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
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

// CreateUser add a new user to database
// INSERT INTO `user` (`created_at`,`updated_at`,`deleted_at`,`username`,`password`,`role`) VALUES ('2020-09-07 12:40:05','2020-09-07 12:40:05',NULL,'panda','123456',0)
func CreateUser(user *User) int {
	user.Password = scryptPassword(user.Password)
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

func scryptPassword(password string) string {
	const keyLen = 10
	salt := []byte{53, 33, 26, 234, 123, 234, 42, 54}
	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, keyLen)

	if err != nil {
		log.Fatal(err.Error())
		return password
	}
	ret := base64.StdEncoding.EncodeToString(dk)
	return ret
}
