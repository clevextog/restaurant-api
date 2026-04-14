package main

import (
	"errors"
	"strings"
)

var validationEmail = []string{"@mail.ru", "@yandex.ru", "@gmail.com", "@yahoo.com", "@rambler.ru"}
var notValidationEmail = []string{"#", "$", "%", "^", "&", "*", "(", ")", "=", "`", "~", "{", "[", "}", "]", ":", ";", "'", ",", "<", ">", "/", "?", "|"}
var notValidationPassword = []string{"#", "$", "%", "^", "&", "*", "(", ")", "=", "`", "~", "{", "[", "}", "]", ":", ";", "'", ",", "<", ">", "/", "?", "|"}

func validEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	dog := "@"
	counter := 0
	for _, symbol := range email {
		if string(symbol) == dog {
			counter++
		}
	}
	if counter != 1 {
		return errors.New("invalid email")
	}
	if strings.HasPrefix(email, "@") {
		return errors.New("invalid email")
	}
	if strings.HasSuffix(email, "@") {
		return errors.New("invalid email")
	}
	for _, forbiddenSymbol := range notValidationEmail {
		if strings.Contains(email, forbiddenSymbol) {
			return errors.New("invalid email")
		}
	}
	for _, domain := range validationEmail {
		if strings.HasSuffix(email, domain) {
			return nil
		}
	}
	return errors.New("invalid email domain")
}

func validPassword(password string) error {
	if password == "" {
		return errors.New("password is required")
	}
	for _, forbiddenSymbol := range notValidationPassword {
		if strings.Contains(password, forbiddenSymbol) {
			return errors.New("invalid password")
		}
	}
	if len(password) < 8 {
		return errors.New("password is too short")
	}
	if len(password) > 50 {
		return errors.New("password is too long")
	}
	digitCounter := 0
	for _, digit := range password {
		if digit >= '0' && digit <= '9' {
			digitCounter++
		}
	}
	if digitCounter < 3 {
		return errors.New("at least 3 digits are required in the password")
	}
	uppercaseLetter := 0
	for _, letter := range password {
		if letter >= 'A' && letter <= 'Z' {
			uppercaseLetter++
		}
	}
	if uppercaseLetter < 1 {
		return errors.New("at least one capital letter is required")
	}
	return nil
}
