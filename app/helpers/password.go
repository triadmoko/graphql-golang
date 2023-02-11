package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// encrypt password
func HashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	passwordHash := string(bytePassword)

	return passwordHash, nil
}

// compare password
func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}
