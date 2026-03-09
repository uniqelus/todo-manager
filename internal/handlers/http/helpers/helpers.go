package helpers

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/render"
)

func DecodeRequest(r *http.Request, request any) error {
	if err := render.DecodeJSON(r.Body, request); err != nil {
		if errors.Is(err, io.EOF) {
			return ErrEmptyRequest
		}

		return ErrFailedToDecodeRequset
	}

	return nil
}
