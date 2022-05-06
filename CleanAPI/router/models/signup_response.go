package models

import (
	"strings"

	"github.com/google/uuid"
)

type signUpModelResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSignUpModelResponse(req SignUpModel) signUpModelResponse {
	return signUpModelResponse{
		Id:       uuid.New().String(),
		Name:     req.Name,
		Email:    req.Email,
		Password: strings.Repeat("*", len(req.Password)),
	}
}

func (r signUpModelResponse) String() string {
	return stringify(r)
}
