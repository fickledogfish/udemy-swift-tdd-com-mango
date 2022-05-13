package models

import (
	"example.com/api/crypt"
	"github.com/google/uuid"
)

type user struct {
	AccessToken  string
	Name         string
	Email        string
	PasswordHash []byte
}

func NewUser(requestModel SignUpModel) (newUser user, err error) {
	passHash, err := crypt.HashPassword(requestModel.Password)
	if err != nil {
		return
	}

	return user{
		AccessToken:  uuid.NewString(),
		Name:         requestModel.Name,
		Email:        requestModel.Email,
		PasswordHash: passHash,
	}, nil
}
