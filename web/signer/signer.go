package signer

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	hash2 "hash"
	"net/url"
	"sort"
	"strings"

	"github.com/duke-git/lancet/v2/convertor"
)

func (c *Client) Sign(params map[string]any, sk string) (sign string) {
	data := FormatURLParam(params)

	var h hash2.Hash
	var hash []byte
	switch c.Algorithm {
	case AlgorithmSha256:
		h = hmac.New(sha256.New, []byte(sk))
	default:
		h = md5.New()
		data = fmt.Sprintf("%s&%s=%s", data, c.KeySK, sk)
	}

	h.Write([]byte(data))
	hash = h.Sum(nil)

	return hex.EncodeToString(hash)
}

func (c *Client) VerifySign(params map[string]any, secretKey string, sign string) bool {
	resign := c.Sign(params, secretKey)

	return sign == resign
}

func FormatURLParam(body map[string]any) (urlParam string) {
	var (
		buf  strings.Builder
		keys []string
	)

	for k := range body {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		v, ok := body[k].(string)
		if !ok {
			v = convertor.ToString(body[k])
		}
		if v != "" {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}

	if buf.Len() <= 0 {
		return ""
	}

	return buf.String()[:buf.Len()-1]
}
