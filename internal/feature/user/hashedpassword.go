package userFeature

import (
	"golang.org/x/crypto/bcrypt"
	"log/slog"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		slog.Error("error generate hash password", err.Error())
		return "", err
	}
	return string(bytes), nil
}
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
