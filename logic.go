package main

import (
	"golang.org/x/crypto/bcrypt"
)

func Register(email, password string) error {
	err := validEmail(email)
	if err != nil {
		return err
	}
	err = validPassword(password)
	if err != nil {
		return err
	}
	_, err = GetUserByEmail(email)
	if err == nil {
		return ErrUserExists
	}
	if err != ErrUserNotFound {
		return err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	passwordHash := string(pass)
	err = CreateUser(email, passwordHash)
	if err != nil {
		return err
	}
	return nil
}
