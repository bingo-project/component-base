package token

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
	ErrTokenInvalid  = errors.New("couldn't handle this token")

	config = Config{}
	once   sync.Once
)

type Config struct {
	SecretKey string
	TTL       uint
}

type CustomClaims struct {
	Info interface{} `json:"info"`
	jwt.RegisteredClaims
}

type Response struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

func Init(secretKey string, ttl uint) {
	once.Do(func() {
		config.SecretKey = secretKey
		config.TTL = ttl
	})
}

// Sign a token by jwt secret.
func Sign(subject string, info interface{}) (*Response, error) {
	// Register claims
	claims := CustomClaims{
		info,
		jwt.RegisteredClaims{
			Subject:   subject,
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.TTL))),
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return nil, err
	}

	resp := &Response{
		AccessToken: tokenString,
		ExpiresAt:   claims.ExpiresAt.Time,
	}

	return resp, err
}

// ParseRequest Parse token from request header.
func ParseRequest(c *gin.Context) (*CustomClaims, error) {
	t := GetBearerToken(c)

	return Parse(t, config.SecretKey)
}

// GetBearerToken Get bearer token from request header.
func GetBearerToken(c *gin.Context) string {
	header := c.Request.Header.Get("Authorization")

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

// Parse token by secret key.
func Parse(tokenString string, key string) (*CustomClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
