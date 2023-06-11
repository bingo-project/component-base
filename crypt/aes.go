package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
)

type AES struct {
	key []byte
}

// NewAES returns a new AES hasher.
func NewAES(key string) *AES {
	keyBytes := []byte(key)
	return &AES{
		key: keyBytes,
	}
}

func (a *AES) EncryptString(value string) (string, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(value)

	iv := make([]byte, 12)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, iv, plaintext, nil)

	jsonEncoded, err := json.Marshal(map[string][]byte{
		"iv":    iv,
		"value": ciphertext,
	})
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(jsonEncoded), nil
}

// DecryptString decrypts the given iv and ciphertext, and returns the plaintext.
func (a *AES) DecryptString(payload string) (string, error) {
	decodePayload, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return "", err
	}

	decodeJson := make(map[string][]byte)
	err = json.Unmarshal(decodePayload, &decodeJson)
	if err != nil {
		return "", err
	}

	// check if the json payload has the correct keys
	if _, ok := decodeJson["iv"]; !ok {
		return "", errors.New("decrypt payload error: missing iv key")
	}
	if _, ok := decodeJson["value"]; !ok {
		return "", errors.New("decrypt payload error: missing value key")
	}

	decodeIv := decodeJson["iv"]
	decodeCiphertext := decodeJson["value"]

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, decodeIv, decodeCiphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
