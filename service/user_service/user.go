package user_service

import (
	"github.com/zhangtao25/mangostreet-ser-gin/models"
)

type User struct {
	Id   int
	Username string
	Password string
	Vcode string
}

func (u *User) Get() (string, error) {
	token, err := models.AuthUsersByVerificationCode(u.Username,u.Vcode)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u *User) ExistByUsername() (bool, error) {
	return models.ExistUserByUsername(u.Username)
}