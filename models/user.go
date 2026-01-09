package models

import (
	"auto-pharmacy/database"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", errors.Join(errors.New("Password hash error"), err)
	}
	return string(hashedPassword), nil
}

type User struct {
	gorm.Model
	Email    string `gorm:"email;unique" json:"email"`
	Password string `gorm:"password" json:"password"`
	Name     string `gorm:"name" json:"name"`
}

func UserMigrate() error {
	return database.MysqlDB.DB.AutoMigrate(&User{})
}

func NewUser(email string, password string, name string) (*User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{Email: email, Password: hashedPassword, Name: name}, nil
}

func (u *User) VerifyPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
