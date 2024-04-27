package token

import (
	"fmt"
	"net/http"
)

// ParseRequest Parse token from request header.
func ParseRequest(req *http.Request) (*CustomClaims, error) {
	t := GetBearerToken(req)

	return Parse(t, config.SecretKey)
}

// GetBearerToken Get bearer token from request header.
func GetBearerToken(req *http.Request) string {
	header := req.Header.Get("Authorization")

	if len(header) == 0 {
		return ""
	}

	var t string

	// Get token from header
	_, err := fmt.Sscanf(header, "Bearer %s", &t)
	if err != nil {
		return ""
	}

	return t
}

// ParseToken Parse given token.
func ParseToken(req *http.Request, token string) (*CustomClaims, error) {
	return Parse(token, config.SecretKey)
}
