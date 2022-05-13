package models

import "encoding/json"

type returnError struct {
	Error string `json:"error"`
}

func NewReturnError(message string) returnError {
	return returnError{
		Error: message,
	}
}

// Implementing encoding.BinaryMarshaler --------------------------------------

func (e returnError) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}
