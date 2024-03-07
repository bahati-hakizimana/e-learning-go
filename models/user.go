package models

import (

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_nsme" validate:"required, min=2, max=100"`
	LastName  string `json:"last_nsme" validate:"required, min=2, max=100"`
	Email     string `json:"email" validate:"email,required"`
	Password  []byte `json:"."`
	Phone     string `json:"phone" validate:"required"`
}

func (user *User) SetPassword(password string){
	hashedPassword,_:=bcrypt.GenerateFromPassword([]byte(password),14)
	user.Password=hashedPassword
}