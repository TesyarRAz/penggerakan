package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes)
}

func CheckPasswordHash(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
