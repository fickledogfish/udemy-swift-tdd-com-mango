package crypt

import "golang.org/x/crypto/bcrypt"

type PasswordComparer interface {
	MatchesHash([]byte) bool
}

type passwordComparer struct {
	password []byte
}

func NewPasswordComparer(password string) PasswordComparer {
	return passwordComparer{
		password: []byte(password),
	}
}

// Implementing PasswordComparer ----------------------------------------------

func (h passwordComparer) MatchesHash(hash []byte) bool {
	return nil == bcrypt.CompareHashAndPassword(hash, h.password)
}
