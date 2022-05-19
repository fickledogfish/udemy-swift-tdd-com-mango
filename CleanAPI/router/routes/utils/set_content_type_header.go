package utils

import (
	"fmt"
	"net/http"

	"example.com/api/log"
)

type headerName string

const (
	contentTypeHeaderName headerName = "Content-Type"
)

type ContentType int

const (
	ContentTypeApplicationJSON ContentType = iota
)

func (c ContentType) String() string {
	switch c {
	case ContentTypeApplicationJSON:
		return "application/json"

	default:
		log.Warn("Unknown ContentType: %d", c)
		return ""
	}
}

func SetContentTypeHeader(w http.ResponseWriter, contentType ContentType) {
	headerName := string(contentTypeHeaderName)
	headerContent := fmt.Sprintf("%s; charset=utf-8", contentType)

	w.Header().Set(headerName, headerContent)
}
