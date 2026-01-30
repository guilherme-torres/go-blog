package utils

import (
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyPasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func Sha256Hash(s string) ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(s)); err != nil {
		return nil, err
	}
	bs := h.Sum(nil)
	return bs, nil
}

func GenerateRandomBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	return buf, nil
}
