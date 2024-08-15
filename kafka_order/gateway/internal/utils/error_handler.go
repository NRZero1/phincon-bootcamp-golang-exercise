package utils

import "errors"

var (
	ErrCreateRequest = errors.New("error when trying to create request")
	ErrClientSend = errors.New("error when trying to send http request")
)