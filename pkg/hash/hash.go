package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func MatchPasswords(inputPassword, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(inputPassword))
	return err == nil
}
