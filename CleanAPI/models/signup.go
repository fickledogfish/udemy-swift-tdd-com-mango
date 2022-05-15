package models

import (
	"bytes"
	"encoding/json"
)

type SignUp struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

// Implementing encoding.BinaryUnmarshaler ------------------------------------

func (s *SignUp) UnmarshalBinary(data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	return decoder.Decode(&s)
}
