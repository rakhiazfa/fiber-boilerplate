package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	ProfilePicture *string `gorm:"type:varchar(255);default:null"`
	Name           string  `gorm:"type:varchar(100)"`
	Username       string  `gorm:"type:varchar(100);unique"`
	Email          string  `gorm:"type:varchar(255);unique"`
	Password       string  `gorm:"type:varchar(255)"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := u.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = hash
	}

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") && u.Password != "" {
		hash, err := u.HashPassword(u.Password)
		if err != nil {
			return err
		}

		u.Password = hash
	}

	return
}

func (User) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
