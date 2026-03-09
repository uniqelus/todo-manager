package helpers

import "errors"

var (
	ErrEmptyRequest          = errors.New("request body is empty")
	ErrFailedToDecodeRequset = errors.New("failed to decode request")
)
