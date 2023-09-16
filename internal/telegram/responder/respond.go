package responder

import (
	"github.com/pkg/errors"
	"net/http"
)

func respondError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrInvalidRequestBody):
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
