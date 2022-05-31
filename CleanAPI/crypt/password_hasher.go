package crypt

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	Hash(string) ([]byte, error)
}

type passwordHasher struct {
	HashCost int
}

func NewPasswordHasher() PasswordHasher {
	return passwordHasher{
		HashCost: bcrypt.DefaultCost,
	}
}

// Implementing PasswordHasher ------------------------------------------------

func (h passwordHasher) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), h.HashCost)
}
