package signer

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_Sign(t *testing.T) {
	sk := "test-sk"
	params := map[string]any{
		"username":  "Peter",
		"action":    "test",
		"datetime":  "2024-01-01 00:00:00",
		"favorites": []string{"apple", "banana"},
	}

	queryString := FormatURLParam(params)

	signer := New()
	signer.SetKeySK(KeySkDefault)

	// Sign by md5
	h := md5.New()
	queryData := fmt.Sprintf("%s&%s=%s", queryString, signer.KeySK, sk)
	h.Write([]byte(queryData))
	hash := h.Sum(nil)
	data := hex.EncodeToString(hash)
	sign := signer.Sign(params, sk)
	assert.Equal(t, sign, data)

	// Sign by sha256
	signer.SetAlgorithm(AlgorithmSha256)
	h = hmac.New(sha256.New, []byte(sk))
	h.Write([]byte(queryString))
	hash = h.Sum(nil)
	data = hex.EncodeToString(hash)
	sign = signer.Sign(params, sk)
	assert.Equal(t, sign, data)
}

func TestClient_VerifySign(t *testing.T) {
	sk := "test-sk"
	params := map[string]any{}

	signer := New()
	sign := signer.Sign(params, sk)
	ret := signer.VerifySign(params, sk, sign)
	assert.True(t, ret)

	params = map[string]any{
		"username":  "Peter",
		"action":    "test",
		"datetime":  "2024-01-01 00:00:00",
		"favorites": []string{"apple", "banana"},
	}
	sign = signer.Sign(params, sk)
	ret = signer.VerifySign(params, sk, sign)
	assert.True(t, ret)
}

func TestFormatURLParam(t *testing.T) {
	var params map[string]any

	// Test empty
	queryString := FormatURLParam(params)
	assert.Empty(t, queryString)

	// Test format
	params = map[string]any{
		"username":  "Peter",
		"action":    "test",
		"datetime":  "2024-01-01 00:00:00",
		"favorites": []string{"apple", "banana"},
	}
	queryString = FormatURLParam(params)
	assert.Equal(t, queryString, "action=test&datetime=2024-01-01+00%3A00%3A00&favorites=%5B%22apple%22%2C%22banana%22%5D&username=Peter")
}
