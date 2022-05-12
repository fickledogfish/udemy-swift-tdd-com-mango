package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type signUpModelResponse struct {
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

func NewSignUpModelResponse(req SignUpModel) signUpModelResponse {
	return signUpModelResponse{
		Name:        req.Name,
		AccessToken: uuid.New().String(),
	}
}

// Implementing encoding.BinaryMarshaler --------------------------------------

func (s signUpModelResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
