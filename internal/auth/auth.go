package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts an API Key from the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("nalformed authentication header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("nalformed first part of authentication header")
	}

	return vals[1], nil
}
