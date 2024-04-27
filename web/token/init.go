package token

import (
	"sync"
)

var (
	config *Client
	once   sync.Once
)

func Init(secretKey string, ttl uint) {
	once.Do(func() {
		config = New(secretKey, ttl)
	})
}

func SetIssuer(issuer string) {
	config.SetIssuer(issuer)
}

// Sign a token by jwt secret.
func Sign(subject string, info any) (*Response, error) {
	return config.Sign(subject, info)
}

// Parse token by secret key.
func Parse(tokenString string, key string) (*CustomClaims, error) {
	return config.Parse(tokenString, key)
}
