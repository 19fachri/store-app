package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password cannot be empty")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", errors.New("error generating password")
	}

	return string(passwordHash), nil
}

func VerifyPassword(password string, hashedPassword string) error {
	bytePassword := []byte(password)
	bytePasswordHash := []byte(hashedPassword)
	return bcrypt.CompareHashAndPassword(bytePasswordHash, bytePassword)
}
