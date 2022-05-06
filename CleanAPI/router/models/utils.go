package models

import (
	"encoding/json"
	"fmt"
)

func stringify(a any) string {
	str, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(str)
}
