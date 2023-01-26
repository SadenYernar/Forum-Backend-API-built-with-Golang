package service

import (
	"forumv2/internal/models"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

const salt = "kjh4kj3b24b2jk43b32jk4b234b"

func generateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func checkPassword(password string) bool {
	if len(password) <= 7 && len(password) > 30 {
		return false
	}

	for _, w := range password {
		if w < 33 || w > 126 {
			return false
		}
	}

	return true
}

func checkName(name string) bool {
	for _, w := range name {
		if (w < 48 || w > 57) && (w < 65 || w > 90) && (w < 97 || w > 122) {
			return false
		} else if len(name) == 0 {
			return false
		}
	}

	return true
}

func userValidation(user models.User) bool {
	if !checkName(user.Name) {
		return false
	} 
	 if !checkPassword(user.Password) {
		return false
	} 
	 if !checkUsername(user.Username) {
		return false
	} 
	 if !checkEmail(user.Email) {
		return false
	}

	return true
}

func checkUsername(username string) bool {
	for _, w := range username {
		if (w < 33 || w > 126) || len(username) == 0 {
			return false
		}
	}

	return true
}

func checkEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
