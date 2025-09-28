package repository

import "fmt"

var users = map[string]string{}

var ErrUserExists = fmt.Errorf("user already exists")

func SaveUser(username, hashedPassword string) error {
	if _, ok := users[username]; ok {
		return ErrUserExists
	}
	users[username] = hashedPassword
	return nil
}

func GetHashedPassword(username string) (string, bool) {
	storedPassword, ok := users[username]
	return storedPassword, ok
}

func UserExists(username string) bool {
	_, ok := users[username]
	return ok
}
