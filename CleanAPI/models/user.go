package models

import (
	"example.com/api/crypt"
	"github.com/google/uuid"
)

type User struct {
	AccessToken  string
	Name         string
	Email        string
	PasswordHash []byte
}

func NewUser(
	passwordHasher crypt.PasswordHasher,
	requestModel SignUp,
) (newUser User, err error) {
	passHash, err := passwordHasher.Hash(requestModel.Password)
	if err != nil {
		return
	}

	return User{
		AccessToken:  uuid.NewString(),
		Name:         requestModel.Name,
		Email:        requestModel.Email,
		PasswordHash: passHash,
	}, nil
}
