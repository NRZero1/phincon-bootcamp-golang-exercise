package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Response struct {
	Message string
	StatusCode int
	Data interface{}
}

func NewResponse(message string, statusCode int, data interface{}) *Response {
	return &Response{
		Message: message,
		StatusCode: statusCode,
		Data: data,
	}
}

func (r *Response) Write(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	res, err := json.Marshal(&r)

	if err != nil {
		log.Fatal().Msg("Error marshaling")
	}

	w.Write(res)
}