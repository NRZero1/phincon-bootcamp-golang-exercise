package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 15)

	return string(hashed), err
}

func CheckPasswordHash(password string, hash string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}