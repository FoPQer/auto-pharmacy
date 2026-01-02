package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, errors.Join(errors.New("Password hash error"), err)
	}
	return hashedPassword, nil
}

type User struct {
	Email    string `gorm:"email"`
	Password []byte `gorm:"password"`
	Name     string `gorm:"name"`
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
