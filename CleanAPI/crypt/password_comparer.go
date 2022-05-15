package crypt

import "golang.org/x/crypto/bcrypt"

type IPasswordComparer interface {
	MatchesHash([]byte) bool
}

type passwordComparer struct {
	password []byte
}

func NewPasswordComparer(password string) passwordComparer {
	return passwordComparer{
		password: []byte(password),
	}
}

// Implementing IPasswordComparer ---------------------------------------------

func (h passwordComparer) MatchesHash(hash []byte) bool {
	return nil == bcrypt.CompareHashAndPassword(hash, h.password)
}
