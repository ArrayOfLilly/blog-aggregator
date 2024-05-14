package auth

import (
	"fmt"
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("not auth header included in request")

func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	fmt.Printf("%v\n", authHeader)
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	fmt.Printf("%v\n", splitAuth)
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	fmt.Printf("%v\n", splitAuth[1])
	return splitAuth[1], nil
}

func CheckApiKey(savedApiKey string, apiKey string) bool {
	fmt.Printf("savedApiKey %v\n", savedApiKey)
	fmt.Printf("apiKey %v\n", apiKey)
	return (savedApiKey == apiKey)
}