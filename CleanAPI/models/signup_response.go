package models

import (
	"encoding/json"
)

type signUpModelResponse struct {
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

func NewSignUpModelResponse(u User) signUpModelResponse {
	return signUpModelResponse{
		Name:        u.Name,
		AccessToken: u.AccessToken,
	}
}

// Implementing encoding.BinaryMarshaler --------------------------------------

func (s signUpModelResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
