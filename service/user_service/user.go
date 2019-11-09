package user_service

import (
	"fmt"
	"github.com/zhangtao25/mangostreet-ser-gin/models"
)

type User struct {
	Id   int
	Username string
	Password string
	Vcode string
}

func (u *User) Get() (*models.User, error) {
	fmt.Print(u.Username)
	token, err := models.AuthUsersByVerificationCode(u.Username,u.Vcode)
	if err != nil {
		return nil, err
	}

	return token, err
}