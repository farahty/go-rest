package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *User) BeforeCreate(tr *gorm.DB) (err error) {
	if u.Password != nil {
		if *u.Password == "" {
			return
		}
		hash, err := MakeHash(*u.Password)
		if err != nil {
			return nil
		}
		u.Password = &hash
	}
	return
}

func MakeHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(password))
	return err == nil
}
