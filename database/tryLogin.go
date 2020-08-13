package database

import (
	"github.com/ryukazari/twinstgo/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin try to login the user using searching in the database
func TryLogin(email, password string) (models.User, bool) {
	usu, encontrado, _ := UserExists(email)
	if !encontrado {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
