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
