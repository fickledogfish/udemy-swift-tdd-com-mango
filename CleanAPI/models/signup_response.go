package models

import (
	"encoding/json"
)

type signUpResponse struct {
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

func NewSignUpResponse(u User) signUpResponse {
	return signUpResponse{
		Name:        u.Name,
		AccessToken: u.AccessToken,
	}
}

// Implementing encoding.BinaryMarshaler --------------------------------------

func (s signUpResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
