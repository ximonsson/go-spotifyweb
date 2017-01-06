package spotifyweb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var client = &http.Client{}

// Make an API request.
// If the current rate limit is reached it waits it out and retries.
func apiCall(path string, params url.Values, method string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s%s?%s", apiURL, apiVersion, path, params.Encode())
	req, e := http.NewRequest(method, url, nil)
	if e != nil {
		return nil, e
	}
	resp, e := client.Do(req)
	// handle status code
	switch resp.StatusCode {
	case http.StatusBadRequest:
		return nil, newBadRequestError(resp)
	// unauthorized
	case http.StatusUnauthorized:
		return nil, newUnauthorizedError(resp)
	// rate limit exceeded
	case http.StatusTooManyRequests:
		return apiCall(path, params, method)
	default:
		return resp, e
	}
}

// Same as doRequest but it will also decode JSON response into res.
func doRequest(path string, params url.Values, method string, res interface{}) error {
	resp, e := apiCall(path, params, method)
	if e != nil {
		return e
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(res)
}
