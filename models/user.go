package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	UserName string `json:"userName" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone" gorm:"unique"`
}

func (u *User) BeforeUpdate(tr *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := MakePassword(u.Password)
		if err != nil {
			return nil
		}
		u.Password = hash
	}
	return
}

func (u *User) BeforeCreate(tr *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := MakePassword(u.Password)
		if err != nil {
			return nil
		}
		u.Password = hash
	}
	return
}

func MakePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
