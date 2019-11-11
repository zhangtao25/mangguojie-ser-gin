package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username string
	Password string
	Vcode string
	Token string
}

//这个模型，传进来用户名验证码，必定返回一个字符串
func AuthUsersByVerificationCode(username,vcode string) ( string , error) {
	var user User
	err := db.Where("username = ? AND vcode = ? ", username, vcode).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}
	if user.Username != "" {
		return user.Token, nil
	}
	return "", nil
}

func ExistUserByUsername(username string) (bool, error) {
	var user User
	err := db.Where("username = ? ", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.Username != "" {
		return true, nil
	}

	return false, nil
}
