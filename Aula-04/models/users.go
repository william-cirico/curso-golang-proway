package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type Users []User

func (us Users) ToJSON() Users {
	for i := range us {
		us[i].Password = ""
	}

	return us
}

func (u User) ToJSON() User {
	u.Password = ""

	return u
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	u.Password = string(bytes)

	return
}

func (u User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
