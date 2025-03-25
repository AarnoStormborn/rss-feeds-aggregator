package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Authorization: ApiKey {apikey}

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(apiKey, " ")
	if len(vals) != 2 {
		return "", errors.New("incorrect header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("incorrect header title")
	}

	return vals[1], nil
}
