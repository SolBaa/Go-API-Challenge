package viewmodels

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInternal        = errors.New("Internal Server Error")
	ErrNotFound        = errors.New("Not Found")
	ErrBadRequest      = errors.New("Bad Request, Malformed Request Body")
	ErrCompanyNotFound = errors.New("Company Not Found")
	ErrUserNotFound    = errors.New("User Not Found")
)

func JSONError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err.Error())
}
