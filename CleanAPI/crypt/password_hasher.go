package crypt

import "golang.org/x/crypto/bcrypt"

type IPasswordHasher interface {
	Hash(string) ([]byte, error)
}

type passwordHasher struct {
	HashCost int
}

func NewPasswordHasher() passwordHasher {
	return passwordHasher{
		HashCost: bcrypt.DefaultCost,
	}
}

// Implementing IPasswordHasher -----------------------------------------------

func (h passwordHasher) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), h.HashCost)
}
