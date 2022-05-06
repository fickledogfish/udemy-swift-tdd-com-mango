package models

type ReturnError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r ReturnError) String() string {
	return stringify(r)
}
