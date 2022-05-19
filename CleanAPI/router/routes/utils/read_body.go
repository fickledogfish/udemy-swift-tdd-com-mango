package utils

import (
	"io"
	"net/http"
)

func ReadBody(req *http.Request, limitBytes int64) ([]byte, error) {
	return io.ReadAll(io.LimitReader(req.Body, limitBytes))
}
