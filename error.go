package spotifyweb

import (
	"encoding/json"
	"net/http"
)

// apiError is the base API error struct
type apiError struct {
	Msg string
}

func (e apiError) Error() string {
	return e.Msg
}

func newError(r *http.Response) error {
	err := struct {
		Error struct {
			Status  string `json:"string"`
			Message string `json:"string"`
		} `json:"error"`
	}{}
	defer r.Body.Close()
	if e := json.NewDecoder(r.Body).Decode(&err); e != nil {
		return e
	}
	return &apiError{err.Error.Message}
}

// ErrorBadRequest is returned in case of 401 errors.
type ErrorBadRequest struct {
	apiError
}

func newBadRequestError(r *http.Response) error {
	e := newError(r)
	if _, ok := e.(apiError); ok {
		return &ErrorBadRequest{e.(apiError)}
	} else {
		return e
	}
}

// ErrorUnauthorized is returned in case of 403 errors.
type ErrorUnauthorized struct {
	apiError
}

func newUnauthorizedError(r *http.Response) error {
	err := struct {
		Error       string `json:"error"`
		Description string `json:"error_description"`
	}{}
	defer r.Body.Close()
	if e := json.NewDecoder(r.Body).Decode(&err); e != nil {
		return e
	}
	return &ErrorUnauthorized{apiError{err.Description}}
}

// ErrorTooManyRequests in case rate limit has exceeded.
type ErrorTooManyRequests struct {
	apiError
}
