package responses

import (
	"encoding"
	"net/http"

	"example.com/api/router/models"
)

func Ok(w http.ResponseWriter, response encoding.BinaryMarshaler) {
	writeResponse(w, http.StatusOK, response)
}

func BadRequest(w http.ResponseWriter, message string) {
	writeResponse(w, http.StatusForbidden, models.NewReturnError(message))
}

func InternalServerError(w http.ResponseWriter, message string) {
	writeResponse(w, http.StatusInternalServerError, models.NewReturnError(message))
}

func Forbidden(w http.ResponseWriter, message string) {
	writeResponse(w, http.StatusForbidden, models.NewReturnError(message))
}

func writeResponse(
	w http.ResponseWriter,
	code int,
	response encoding.BinaryMarshaler,
) {
	encodedMessage, err := response.MarshalBinary()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\": \"Internal server error\"}"))
	}

	w.WriteHeader(code)
	w.Write(encodedMessage)
}
