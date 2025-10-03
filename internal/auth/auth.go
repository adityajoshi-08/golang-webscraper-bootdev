package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Example: Authorization: APIKey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization header provided")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "APIKey" {
		return "", errors.New("invalid authorization header format")
	}
	return vals[1], nil
}
