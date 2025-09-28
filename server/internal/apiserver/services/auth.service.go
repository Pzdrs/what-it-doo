package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"pycrs.cz/what-it-do/internal/apiserver/repository"
)


func RegisterUser(username, password string) (string, error) {
	if repository.UserExists(username) {
		return "", fmt.Errorf("user already exists")
	}	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	repository.SaveUser(username, string(hashedPassword))
	return "Registration successful", nil
}

func AuthenticateUser(username, password string) bool {
	if storedPassword, ok := repository.GetHashedPassword(username); ok {
		if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err == nil {
			return true
		}
	}
	return false
}