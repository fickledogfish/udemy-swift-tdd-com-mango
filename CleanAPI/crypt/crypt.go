package crypt

import "golang.org/x/crypto/bcrypt"

const (
	hashCost = bcrypt.DefaultCost
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), hashCost)
}

func HashMatchesPassword(hash []byte, password string) bool {
	return nil == bcrypt.CompareHashAndPassword(hash, []byte(password))
}
