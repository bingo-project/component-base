package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var ErrTokenInvalid = errors.New("couldn't handle this token")

type Client struct {
	SecretKey string
	TTL       uint
	Issuer    string
}

type CustomClaims struct {
	Info any `json:"info"`
	jwt.RegisteredClaims
}

type Response struct {
	AccessToken string    `json:"accessToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

func New(secretKey string, ttl uint) *Client {
	return &Client{
		SecretKey: secretKey,
		TTL:       ttl,
	}
}

func (client *Client) SetIssuer(issuer string) *Client {
	client.Issuer = issuer

	return client
}

// Sign a token by jwt secret.
func (client *Client) Sign(subject string, info any) (*Response, error) {
	// Register claims
	claims := CustomClaims{
		info,
		jwt.RegisteredClaims{
			Issuer:    client.Issuer,
			Subject:   subject,
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(client.TTL))),
		},
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(client.SecretKey))
	if err != nil {
		return nil, err
	}

	resp := &Response{
		AccessToken: tokenString,
		ExpiresAt:   claims.ExpiresAt.Time,
	}

	return resp, err
}

// Parse token by secret key.
func (client *Client) Parse(tokenString string) (*CustomClaims, error) {
	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(client.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}
