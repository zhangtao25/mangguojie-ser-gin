package models

import "fmt"

type User struct {
	Model
	Username string
	Password string
	Vcode string
}


func AuthUsersByVerificationCode(username,vcode string) ( *User , error) {
	fmt.Print(username, vcode)
	fmt.Print(11111)
	//var note Note
	var user User
	db.Where("username = ?", username).Find(&user)


	return &user, nil
}